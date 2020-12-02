package xojo

import (
	"fmt"
	"log"

	"github.com/joseluisq/goipcc"
)

// ProjectCommands defines a command for run a Xojo project.
type ProjectCommands struct {
	sock *goipcc.IPCSockClient
}

// Open opens a specific Xojo project.
func (c *ProjectCommands) Open(xojoProjectFilePath string, handler func(data []byte, err error)) error {
	str := fmt.Sprintf("{\"tag\":\"build\",\"script\":\"OpenFile(\\\"%s\\\")\nprint \\\"Project is opened.\\\"\"}\x00", xojoProjectFilePath)
	log.Println("open project command sent:", str)
	_, err := c.sock.Write([]byte(str), handler)
	if err != nil {
		return err
	}
	return nil
}

// Run runs the current opened Xojo project.
func (c *ProjectCommands) Run(handler func(data []byte, err error)) error {
	str := "{\"tag\":\"build\",\"script\":\"DoCommand(\\\"RunApp\\\")\nprint \\\"App is running.\\\"\"}\x00"
	log.Println("run project command sent:", str)
	_, err := c.sock.Write([]byte(str), handler)
	if err != nil {
		return err
	}
	return nil
}

// Close closes the current opened project.
func (c *ProjectCommands) Close(handler func(data []byte, err error)) error {
	str := "{\"tag\":\"build\",\"script\":\"CloseProject(False)\nprint \\\"Default app closed.\\\"\"}\x00"
	log.Println("close project command sent:", str)
	_, err := c.sock.Write([]byte(str), handler)
	if err != nil {
		return err
	}
	return nil
}

// BuildOptions defines build Xojo project options.
type BuildOptions struct {
	// Target operating system
	OS string
	// Target architecture
	Arch string
	// If reveal is true then the built app is displayed using the OS file manager.
	Reveal bool
}

// Build builds current opened Xojo project.
func (c *ProjectCommands) Build(opt BuildOptions, handler func(data []byte, err error)) error {
	var buildType int

	// IDE Scripting building commands
	// https://docs.xojo.com/UserGuide:IDE_Scripting_Building_Commands
	// Value	Build Target		32/64-bit	Architecture
	// 3 		Windows 			32-bit		Intel
	// 4 		Linux				32-bit		Intel
	// 9 		macOS Universal		64-bit		Intel & ARM
	// 12 		Xojo Cloud			32-bit		Intel
	// 14 		iOS Simulator		64-bit		Intel
	// 15 		iOS					64-bit		ARM
	// 16 		Mac (all)			64-bit		Intel
	// 17 		Linux				64-bit		Intel
	// 18 		Linux				32-bit		ARM
	// 19 		Windows				64-bit		Intel
	// 24 		macOS				64-bit		ARM

	if opt.OS == "windows" && opt.Arch == "i386" {
		buildType = 3
	}
	if opt.OS == "windows" && opt.Arch == "amd64" {
		buildType = 19
	}
	if opt.OS == "darwin" && opt.Arch == "amd64" {
		buildType = 16
	}
	if opt.OS == "darwin" && opt.Arch == "arm64" {
		buildType = 24
	}
	if opt.OS == "linux" && opt.Arch == "i386" {
		buildType = 4
	}
	if opt.OS == "linux" && opt.Arch == "amd64" {
		buildType = 17
	}
	if opt.OS == "ios" && opt.Arch == "amd64" {
		buildType = 14
	}
	if opt.OS == "ios" && opt.Arch == "arm64" {
		buildType = 15
	}

	if buildType == 0 {
		return fmt.Errorf("Xojo build options provided were not specified or unsupported")
	}

	var reveal string
	if opt.Reveal {
		reveal = "True"
	} else {
		reveal = "False"
	}

	str := fmt.Sprintf("{\"script\":\"Print BuildApp(%d,%s)\", \"tag\":\"build\"}\x00", buildType, reveal)
	log.Printf("build project options chosen: %s/%s\n", opt.OS, opt.Arch)
	log.Printf("build project command sent: %s\n", str)
	_, err := c.sock.Write([]byte(str), handler)
	if err != nil {
		return err
	}
	return nil
}
