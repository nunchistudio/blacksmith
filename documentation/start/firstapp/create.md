---
title: Creating an application
enterprise: false
---

# Creating an application

## Generate a new application

The best way to create a new Blacksmith application is by using the `generate`
command of the CLI. The following command generates all the required files in the
current directory:
```bash
$ blacksmith generate application --name myapp

```

If you prefer, you can generate a new application inside a directory with the
`--path` flag:
```bash
$ blacksmith generate application --name myapp --path ./myapp

```

The directory will be created if it does not exist yet.

## Tidy Go modules

Blacksmith leverages Go modules for managing dependencies. Before continuing, make
sure to validate and lock the dependencies by executing the command:
```bash
$ cd ./myapp
$ go mod tidy

```
