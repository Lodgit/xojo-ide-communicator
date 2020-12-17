package main

import (
	"log"
	"os"
	"xojoidecom/cmd"
)

// Build-time application values
var (
	versionNumber string = "devel"
	buildTime     string
)

func main() {
	if err := cmd.Execute(os.Args, versionNumber, buildTime); err != nil {
		log.Fatalln(err)
	}
}
