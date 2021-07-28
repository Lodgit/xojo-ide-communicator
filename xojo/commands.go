package xojo

import (
	"fmt"
	"log"
	"strings"

	"github.com/joseluisq/gonetc"
)

// ProjectCommands defines a command for run a Xojo project.
type ProjectCommands struct {
	sock *gonetc.NetClient
}

// Open opens a specific Xojo project.
func (c *ProjectCommands) Open(xojoProjectFilePath string, handler func(data []byte, err error)) (err error) {
	str := fmt.Sprintf("{\"tag\":\"build\",\"script\":\"OpenFile(\\\"%s\\\")\\nprint \\\"Project is opened.\\\"\"}%s", xojoProjectFilePath, XojoNullChar)
	log.Println("open project command sent:", str)
	_, err = c.sock.Write([]byte(str), func(data []byte, err error, done func()) {
		handler(CheckForErrorResponse(data, err))
		done()
	})
	return
}

// Run runs the current opened Xojo project.
func (c *ProjectCommands) Run(handler func(data []byte, err error)) (err error) {
	str := "{\"tag\":\"build\",\"script\":\"DoCommand(\\\"RunApp\\\")\\nprint \\\"App is running.\\\"\"}" + XojoNullChar
	log.Println("run project command sent:", str)
	_, err = c.sock.Write([]byte(str), func(data []byte, err error, done func()) {
		handler(CheckForErrorResponse(data, err))
		done()
	})
	return
}

// Close closes the current opened project.
func (c *ProjectCommands) Close(handler func(data []byte, err error)) (err error) {
	str := "{\"tag\":\"build\",\"script\":\"CloseProject(False)\\nprint \\\"Default app closed.\\\"\"}" + XojoNullChar
	log.Println("close project command sent:", str)
	_, err = c.sock.Write([]byte(str), func(data []byte, err error, done func()) {
		handler(CheckForErrorResponse(data, err))
		done()
	})
	return
}

// BuildOptions defines build Xojo project options.
type BuildOptions struct {
	// Operating system and architecture target pair. E.g linux-amd64
	Target string
	// If reveal is true then the built app is displayed using the OS file manager.
	Reveal bool
}

// Build builds current opened Xojo project.
func (c *ProjectCommands) Build(opt BuildOptions, handler func(data []byte, err error)) (err error) {
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

	target := strings.Split(opt.Target, "-")
	if len(target) == 0 {
		err = fmt.Errorf("one build target was not provided or empty")
		return
	}
	if len(target) < 2 || len(target) > 2 {
		err = fmt.Errorf("one build target has not valid `os-arch` pair value")
		return
	}
	vos := strings.TrimSpace(target[0])
	varch := strings.TrimSpace(target[1])

	var buildType int
	if vos == "linux" && varch == "386" {
		buildType = 4
	}
	if vos == "linux" && varch == "amd64" {
		buildType = 17
	}
	if vos == "darwin" && varch == "amd64" {
		buildType = 16
	}
	if vos == "darwin" && varch == "arm64" {
		buildType = 24
	}
	if vos == "windows" && varch == "386" {
		buildType = 3
	}
	if vos == "windows" && varch == "amd64" {
		buildType = 19
	}
	if vos == "ios" && varch == "amd64" {
		buildType = 14
	}
	if vos == "ios" && varch == "arm64" {
		buildType = 15
	}
	if buildType == 0 {
		err = fmt.Errorf("build target `%s-%s` is not supported by Xojo", vos, varch)
		return
	}

	var reveal string
	if opt.Reveal {
		reveal = "True"
	} else {
		reveal = "False"
	}

	str := fmt.Sprintf("{\"script\":\"Print BuildApp(%d,%s)\", \"tag\":\"build\"}%s", buildType, reveal, XojoNullChar)
	log.Printf("build project options chosen: %s/%s\n", vos, varch)
	log.Printf("build project command sent: %s\n", str)
	_, err = c.sock.Write([]byte(str), func(data []byte, err error, done func()) {
		handler(CheckForErrorResponse(data, err))
		done()
	})
	return
}
