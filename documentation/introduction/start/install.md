---
title: Installation
enterprise: false
---

# Installation

## Installation of the CLI

Blacksmith is available as a single binary for several operating systems, as well
as Docker images for running applications. You can download it from the dedicated
[downloads page](/blacksmith/downloads).

After downloading Blacksmith, you need to unzip the package in an appropriate
directory. Make sure that the `blacksmith` binary is available on your `PATH`,
before continuing.

You can check the locations available on your path by running:
```
$ echo $PATH
```

To verify Blacksmith was installed correctly, try the `blacksmith` command:
```bash
$ blacksmith version
```

You should see the version installed, similar to the following:
```bash
Blacksmith Standard Edition v0.14.0
Built with Go v1.15.5 for darwin/amd64
```

## Installation of the Go package

Now that the command-line is successfully installed, you also need to install the
Go library so you can interact with the library when generating files:
```bash
$ go get github.com/nunchistudio/blacksmith
```

We are ready to create our first data platform!
