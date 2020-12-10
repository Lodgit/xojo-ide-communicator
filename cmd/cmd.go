package cmd

import (
	"log"
	"os"

	cli "github.com/joseluisq/cline"
)

// Execute is the main entry point of the current application.
func Execute(versionNumber string, buildTime string) {
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
		TestCmd(),
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
