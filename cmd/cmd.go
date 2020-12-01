package cmd

import (
	"log"
	"os"
	"xojoidecom/xojo"

	cli "github.com/joseluisq/cline"
)

// Build-time application values
var (
	versionNumber string = "devel"
	buildTime     string
)

// Execute is the main entry point of the current application.
func Execute() {
	app := cli.New()
	app.Name = "xojo-ide-com"
	app.Summary = "CLI client to communicate transparently with Xojo IDE using The Xojo IDE Communication Protocol v2."
	app.Version = versionNumber
	app.BuildTime = buildTime
	app.Commands = []cli.Cmd{
		{
			Name:    "run",
			Summary: "Runs a Xojo opened project in debug mode.",
			Handler: func(ctx *cli.CmdContext) error {
				xo := xojo.New()
				if err := xo.Connect(); err != nil {
					return err
				}
				defer xo.Close()
				err := xo.Commands.Run(func(data []byte, err error) {
					if err != nil {
						log.Fatalln(err)
					}
					log.Println("data received:", string(data))
				})
				return err
			},
		},
		{
			Name:    "build",
			Summary: "Builds a Xojo opened project.",
			Handler: func(ctx *cli.CmdContext) error {
				xo := xojo.New()
				if err := xo.Connect(); err != nil {
					return err
				}
				defer xo.Close()
				err := xo.Commands.Build(func(data []byte, err error) {
					if err != nil {
						log.Fatalln(err)
					}
					log.Println("data received:", string(data))
				})
				return err
			},
		},
		{
			Name:    "close",
			Summary: "Closes a Xojo opened project.",
			Handler: func(ctx *cli.CmdContext) error {
				xo := xojo.New()
				if err := xo.Connect(); err != nil {
					return err
				}
				defer xo.Close()
				err := xo.Commands.Close(func(data []byte, err error) {
					if err != nil {
						log.Fatalln(err)
					}
					log.Println("data received:", string(data))
				})
				return err
			},
		},
	}
	app.Handler = appHandler
	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}

func appHandler(ctx *cli.AppContext) error {
	return nil
}
