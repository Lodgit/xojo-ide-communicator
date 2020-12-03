package cmd

import (
	"log"
	"os"
	"path"
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
	app.Flags = []cli.Flag{
		cli.FlagBool{
			Name:    "use-current-workdir",
			Summary: "Use the current working directory as the base path for \"PROJECT_FILE_PATH\" argument on `run` and `build` commands.",
			Aliases: []string{"w"},
		},
	}
	app.Commands = []cli.Cmd{
		{
			Name:    "run",
			Summary: "Runs a Xojo project in debug mode. Example: xojo-ide-com run [OPTIONS] PROJECT_FILE_PATH",
			Handler: func(ctx *cli.CmdContext) error {
				// 0. Check for project file path argument
				if len(ctx.TailArgs) == 0 || ctx.TailArgs[0] == "" {
					log.Fatalln("xojo project file path was not provided.")
				}
				// Capture the file path argument and check for a "current working directory" usage
				filePath := ctx.TailArgs[0]
				useWorkdir := ctx.AppContext.Flags.IsProvidedFlag("use-current-workdir")
				if useWorkdir {
					cwd, err := os.Getwd()
					if err != nil {
						return err
					}
					filePath = path.Join(cwd, filePath)
				}
				// 1. Xojo socket connection
				xo := xojo.New()
				if err := xo.Connect(); err != nil {
					return err
				}
				defer xo.Close()
				// 2. Close current project first
				err := xo.ProjectCmds.Close(func(data []byte, err error) {
					if err != nil {
						log.Fatalln(err)
					}
					log.Println("data received:", string(data))
				})
				// 3. Open a new specific project
				err = xo.ProjectCmds.Open(filePath, func(data []byte, err error) {
					if err != nil {
						log.Fatalln(err)
					}
					log.Println("data received:", string(data))
				})
				if err != nil {
					log.Fatalln(err)
				}
				// 4. Run current specified project
				err = xo.ProjectCmds.Run(func(data []byte, err error) {
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
			Summary: "Builds a Xojo project. Example: xojo-ide-com build [OPTIONS] PROJECT_FILE_PATH",
			Flags: []cli.Flag{
				cli.FlagStringSlice{
					Name:    "os",
					Summary: "Operating system target(s) such as `linux`, `darwin`, `windows` and `ios`. For multiple targets use a coma-separated list.",
				},
				cli.FlagString{
					Name:    "arch",
					Summary: "Target architecture such as `i386`, `amd64` and `arm64`.",
				},
				cli.FlagBool{
					Name:    "reveal",
					Value:   false,
					Summary: "Open the built application directory using the operating system file manager available.",
				},
			},
			Handler: func(ctx *cli.CmdContext) error {
				// 0. Check for project file path argument
				if len(ctx.TailArgs) == 0 || ctx.TailArgs[0] == "" {
					log.Fatalln("xojo project file path was not provided.")
				}
				// Capture the file path argument and check for a "current working directory" usage
				filePath := ctx.TailArgs[0]
				useWorkdir := ctx.AppContext.Flags.IsProvidedFlag("use-current-workdir")
				if useWorkdir {
					cwd, err := os.Getwd()
					if err != nil {
						return err
					}
					filePath = path.Join(cwd, filePath)
				}
				// 1. Validate arguments
				reveal, err := ctx.Flags.Bool("reveal")
				if err != nil {
					return err
				}
				osStrSlice := ctx.Flags.StringSlice("os")
				if len(osStrSlice) == 0 {
					log.Fatalln("no operating system was specified")
				}
				archStr := ctx.Flags.String("arch")
				if archStr == "" {
					log.Fatalln("no architecture was specified")
				}
				// 2. Xojo socket connection
				xo := xojo.New()
				if err := xo.Connect(); err != nil {
					return err
				}
				defer xo.Close()
				// 3. Close current project first
				err = xo.ProjectCmds.Close(func(data []byte, err error) {
					if err != nil {
						log.Fatalln(err)
					}
					log.Println("data received:", string(data))
				})
				// 4. Open the specified project
				err = xo.ProjectCmds.Open(ctx.TailArgs[0], func(data []byte, err error) {
					if err != nil {
						log.Fatalln(err)
					}
					log.Println("data received:", string(data))
				})
				if err != nil {
					log.Fatalln(err)
				}
				// 5. Build the specified project for each operating system(s) chosen
				for _, osStr := range osStrSlice {
					opts := xojo.BuildOptions{
						OS:     osStr,
						Arch:   archStr,
						Reveal: reveal,
					}
					err = xo.ProjectCmds.Build(opts, func(data []byte, err error) {
						if err != nil {
							log.Fatalln(err)
						}
						log.Println("data received:", string(data))
					})
					if err != nil {
						return err
					}
				}
				// 6. Close current project
				err = xo.ProjectCmds.Close(func(data []byte, err error) {
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
