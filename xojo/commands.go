package xojo

import (
	"bytes"
	"fmt"
	"log"

	"github.com/joseluisq/gonetc"
)

var xojoErrors = [][]byte{
	[]byte("buildError"),
	[]byte("loadError"),
	[]byte("openErrors"),
	[]byte("scriptError"),
}

// checkForErrorResponse just checks many possible errors during commands execution.
func checkForErrorResponse(jsonb []byte, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	for _, e := range xojoErrors {
		if bytes.Contains(jsonb, e) {
			return jsonb, fmt.Errorf(
				"Xojo IDE commands execution error occurred, please check the response output",
			)
		}
	}
	return jsonb, nil
}

// ProjectCommands defines a command for run a Xojo project.
type ProjectCommands struct {
	sock *gonetc.NetClient
}

// Open opens a specific Xojo project.
func (c *ProjectCommands) Open(xojoProjectFilePath string, handler func(data []byte, err error)) error {
	str := fmt.Sprintf("{\"tag\":\"build\",\"script\":\"OpenFile(\\\"%s\\\")\nprint \\\"Project is opened.\\\"\"}%s", xojoProjectFilePath, XojoNullChar)
	log.Println("open project command sent:", str)
	_, err := c.sock.Write([]byte(str), func(data []byte, err error, done func()) {
		handler(checkForErrorResponse(data, err))
		done()
	})
	return err
}

// Run runs the current opened Xojo project.
func (c *ProjectCommands) Run(handler func(data []byte, err error)) error {
	str := "{\"tag\":\"build\",\"script\":\"DoCommand(\\\"RunApp\\\")\nprint \\\"App is running.\\\"\"}" + XojoNullChar
	log.Println("run project command sent:", str)
	_, err := c.sock.Write([]byte(str), func(data []byte, err error, done func()) {
		handler(checkForErrorResponse(data, err))
		done()
	})
	return err
}

// Close closes the current opened project.
func (c *ProjectCommands) Close(handler func(data []byte, err error)) error {
	str := "{\"tag\":\"build\",\"script\":\"CloseProject(False)\nprint \\\"Default app closed.\\\"\"}" + XojoNullChar
	log.Println("close project command sent:", str)
	_, err := c.sock.Write([]byte(str), func(data []byte, err error, done func()) {
		handler(checkForErrorResponse(data, err))
		done()
	})
	return err
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
	// IDE Scripting building options
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

	var buildType int
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

	str := fmt.Sprintf("{\"script\":\"Print BuildApp(%d,%s)\", \"tag\":\"build\"}%s", buildType, reveal, XojoNullChar)
	log.Printf("build project options chosen: %s/%s\n", opt.OS, opt.Arch)
	log.Printf("build project command sent: %s\n", str)
	_, err := c.sock.Write([]byte(str), func(data []byte, err error, done func()) {
		handler(checkForErrorResponse(data, err))
		done()
	})
	return err
}
