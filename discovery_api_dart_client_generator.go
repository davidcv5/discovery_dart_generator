package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

const help string = `
$ discovery_dart_generator -h
Usage:

The discovery generator downloads the discovery documents
and generates an API package. It takes the following options:

-m, --mode                   m=package (creates new package, default)
                             m=files   (update existing package)

Package Mode:

-u, --url                    URL of the discovery documents. 
                             (required)

-o, --output-dir             Output directory of the generated API package.
                             (defaults to "googleapis")

-p, --package-name           Name of the generated API package.
                             (defaults to "googleapis")

-v, --package-version        Version of the generated API package.
                             (defaults to "0.1.0-dev")

-d, --package-description    Description of the generated API package.
                             (defaults to "Auto-generated client libraries.")

-a, --package-author         Author of the generated API package.

-h, --package-homepage       Homepage of the generated API package.

Files Mode:

-o, --output-dir             Output directory of the generated API package.
                             (defaults to "googleapis")

-up, --update-pubspec        Update the pubspec.yaml file with required dependencies. 
                             This will remove comments and might change the layout of the pubspec.yaml file.
                             (defaults to "false")
`

func createStringFlag(name string, short_name string, def string, desc string) *string {
	flagPtr := flag.String(name, def, desc)
	flag.StringVar(flagPtr, short_name, def, desc)
	return flagPtr
}

func createBooleanFlag(name string, short_name string, def bool, desc string) *bool {
	flagPtr := flag.Bool(name, def, desc)
	flag.BoolVar(flagPtr, short_name, def, desc)
	return flagPtr
}

func printHelp() {
	fmt.Println(help)
	os.Exit(2)
}

func cleanAndExit(msg string, err error) {
	if msg != "" {
		fmt.Println(msg)
	}
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Cleanning up...")
	cmd := exec.Command("rm", "-rf", "discoveryapis_generator", "googleapis-discovery-documents")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error cleaning up.")
		fmt.Println("error:", err)
		os.Exit(2)
	}
	os.Exit(0)
}

func main() {

	urlPtr := createStringFlag("url", "u", "", "Discorvery URL")
	modePtr := createStringFlag("mode", "m", "package", "Create (package) or update (files) a package.")
	outputPtr := createStringFlag("output", "o", "googleapis", "Name of the generated API package.")
	packagePtr := createStringFlag("package-name", "p", "googleapis", "Name of the generated API package.")
	versionPtr := createStringFlag("package-version", "v", "0.1.0-dev", "Version of the generated API package.")
	descriptionPtr := createStringFlag("package-description", "d", "Auto-generated client libraries.", "Description of the generated API package.")
	authorPtr := createStringFlag("package-author", "a", "", "Author of generated API package.")
	homepagePtr := createStringFlag("package-homepage", "h", "", "Homepage of the generated API package.")
	updatePtr := createBooleanFlag("update-pubspec", "up", false, "Update the pubspec.yaml file with required dependencies.")

	flag.Usage = printHelp
	flag.Parse()

	if (*modePtr != "package" && *modePtr != "files") || *urlPtr == "" {
		printHelp()
	}

	fmt.Println("\n**********")
	fmt.Println("Starting generator...")
	fmt.Printf("URL: %s\n", *urlPtr)
	fmt.Printf("Package: %s\n", *packagePtr)
	fmt.Printf("Version: %s\n", *versionPtr)
	fmt.Printf("Description: %s\n", *descriptionPtr)
	if *authorPtr != "" {
		fmt.Printf("Author: %s\n", *authorPtr)
	}
	if *homepagePtr != "" {
		fmt.Printf("Homepage: %s\n", *homepagePtr)
	}
	fmt.Println("")

	cmd := exec.Command("mkdir", "googleapis-discovery-documents")
	err := cmd.Run()
	if err != nil {
		cleanAndExit("Error creating directory. Make sure you have write permission.", err)
	}

	fmt.Println("Downloading discovery documents...")
	cmd = exec.Command("curl", "-s", "-o", "googleapis-discovery-documents/generated.json", *urlPtr)
	err = cmd.Run()
	if err != nil {
		cleanAndExit("Error downloading discovery documents.\nMake sure you're using the valid URL", err)
	}

	fmt.Println("Getting dart generator libraries...")
	cmd = exec.Command("git", "clone", "https://github.com/dart-lang/discoveryapis_generator.git")
	err = cmd.Run()
	if err != nil {
		cleanAndExit("Error downloading the dart discovery api generator library.\nURL might have change, please file a bug.", err)
	}

	fmt.Println("Building generator libraries...")
	cmd = exec.Command("pub", "get")
	cmd.Dir = "discoveryapis_generator"
	err = cmd.Run()
	if err != nil {
		cleanAndExit("Error building dart libraries.\nMake sure you have Dart installed and can you can run 'pub' on the commadn line", err)
	}

	fmt.Println("Generating dart client library files...")
	if *modePtr == "package" {
		cmd = exec.Command(
			"bin/generate.dart",
			*modePtr,
			fmt.Sprintf("--package-name=%s", *packagePtr),
			fmt.Sprintf("--package-version=%s", *versionPtr),
			fmt.Sprintf("--package-description=%s", *descriptionPtr),
			"--input-dir=../googleapis-discovery-documents",
			fmt.Sprintf("--output-dir=../%s", *outputPtr))
		if *authorPtr != "" {
			cmd.Args = append(cmd.Args, fmt.Sprintf("--package-author=%s", *authorPtr))
		}
		if *homepagePtr != "" {
			cmd.Args = append(cmd.Args, fmt.Sprintf("--package-homepage=%s", *homepagePtr))
		}
	} else {
		cmd = exec.Command(
			"bin/generate.dart",
			*modePtr,
			"--input-dir=../googleapis-discovery-documents",
			fmt.Sprintf("--output-dir=../%s", *outputPtr),
			fmt.Sprintf("--update-pubspec=%s", *updatePtr))
	}
	cmd.Dir = "discoveryapis_generator"
	err = cmd.Run()
	if err != nil {
		cleanAndExit("Error generating client libraries.", err)
	}
	cleanAndExit("", nil)
}
