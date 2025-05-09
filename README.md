# JoRepos

A collection of Go packages with customized implementations of popular libraries.

## Overview

This repository contains customized wrappers for various Go packages that I frequently use in my projects. Each package has its own `go.mod` file, allowing them to be imported individually.

## Packages

- **log**: Custom logging implementation based on [charmbracelet/log](https://github.com/charmbracelet/log)
- **onepass**: Custom 1Password integration
- **chromedp**: Custom ChromeDP wrapper for browser automation

## Usage

Import the packages directly from GitHub:

```go
import (
    "github.com/jo7e/jorepos/log"
    "github.com/jo7e/jorepos/onepass"
    "github.com/jo7e/jorepos/chromedp"
)
```
