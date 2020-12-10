package main

import (
	"xojoidecom/cmd"
)

// Build-time application values
var (
	versionNumber string = "devel"
	buildTime     string
)

func main() {
	cmd.Execute(versionNumber, buildTime)
}
