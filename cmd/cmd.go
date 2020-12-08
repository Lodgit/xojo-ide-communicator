package cmd

import (
	"log"
	"os"

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
	app.Flags = []cli.Flag{
		cli.FlagBool{
			Name:    "use-current-workdir",
			Summary: "Use the current working directory as the base path for \"PROJECT_FILE_PATH\" argument on `run` and `build` commands.",
			Aliases: []string{"w"},
		},
	}
	app.Commands = []cli.Cmd{
		RunCmd(),
		BuildCmd(),
	}
	app.Handler = appHandler
	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}

func appHandler(ctx *cli.AppContext) error {
	return nil
}
