# discovery_dart_generator

## Description

GO application to generate Dart API Client Libraries based on discovery 
documents.

This package is a wrapper for [discoveryapis_generator] (https://github.com/dart-lang/discoveryapis_generator) to automate the process of generating the Dart API Client Libraries.

## Install

Use [goapp tool][goapp] from the Google App Engine SDK for Go to get the package:

```
GO_APPENGINE/goapp get github.com/davidcv5/discovery_dart_generator
```

If you'll ever need to pull updates from the upstream, execute `git pull`
from the root of this repo.

Alternatively, if you don't have `goapp` for some reason, do the standard

```
go get github.com/davidcv5/discovery_dart_generator
```

If this is not the first time you're "getting" the package,
add `-u` param to get an updated version, i.e. `go get -u ...`.


## Usage

```
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
```