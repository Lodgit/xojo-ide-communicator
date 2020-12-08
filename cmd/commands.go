package cmd

import (
	"log"
	"os"
	"path"
	"time"
	"xojoidecom/xojo"

	cli "github.com/joseluisq/cline"
)

// RunCmd defines the Xojo project `run` command.
func RunCmd() cli.Cmd {
	return cli.Cmd{
		Name:    "run",
		Summary: "Runs a Xojo project in debug mode. Example: xojo-ide-com run [OPTIONS] PROJECT_FILE_PATH",
		Flags: []cli.Flag{
			cli.FlagInt{
				Name:    "delay",
				Summary: "Workaround delay in seconds to wait until the current application is displayed on screen.",
				Value:   5,
				Aliases: []string{"d"},
			},
		},
		Handler: func(ctx *cli.CmdContext) error {
			delay, err := ctx.Flags.Int("delay").Value()
			if err != nil {
				return err
			}
			// 0. Check for project file path argument
			if len(ctx.TailArgs) == 0 || ctx.TailArgs[0] == "" {
				log.Fatalln("xojo project file path was not provided.")
			}
			// Capture the file path argument and check for a "current working directory" usage
			filePath := ctx.TailArgs[0]
			useWorkdir := ctx.AppContext.Flags.Bool("use-current-workdir").IsProvided()
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
			err = xo.ProjectCmds.Close(func(data []byte, err error) {
				if err != nil {
					log.Println(err)
					log.Fatalln(string(data))
				}
				log.Println("data received:", string(data))
			})
			// 3. Open a new specific project
			err = xo.ProjectCmds.Open(filePath, func(data []byte, err error) {
				if err != nil {
					log.Println(err)
					log.Fatalln(string(data))
				}
				log.Println("data received:", string(data))
			})
			if err != nil {
				log.Fatalln(err)
			}
			// 4. Run current specified project
			err = xo.ProjectCmds.Run(func(data []byte, err error) {
				if err != nil {
					log.Println(err)
					log.Fatalln(string(data))
				}
				log.Println("data received:", string(data))
			})
			log.Printf("waiting %d second(s) for the application...\n", delay)
			time.Sleep(time.Duration(delay) * time.Second)
			return err
		},
	}
}

// BuildCmd defines the Xojo project `build` command.
func BuildCmd() cli.Cmd {
	return cli.Cmd{
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
			useWorkdir := ctx.AppContext.Flags.Bool("use-current-workdir").IsProvided()
			if useWorkdir {
				cwd, err := os.Getwd()
				if err != nil {
					return err
				}
				filePath = path.Join(cwd, filePath)
			}
			// 1. Validate arguments
			reveal, err := ctx.Flags.Bool("reveal").Value()
			if err != nil {
				return err
			}
			osStrSlice := ctx.Flags.StringSlice("os").Value()
			if len(osStrSlice) == 0 {
				log.Fatalln("no operating system was specified")
			}
			archStr := ctx.Flags.String("arch").Value()
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
					log.Println(err)
					log.Fatalln(string(data))
				}
				log.Println("data received:", string(data))
			})
			// 4. Open the specified project
			err = xo.ProjectCmds.Open(ctx.TailArgs[0], func(data []byte, err error) {
				if err != nil {
					log.Println(err)
					log.Fatalln(string(data))
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
						log.Println(err)
						log.Fatalln(string(data))
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
					log.Println(err)
					log.Fatalln(string(data))
				}
				log.Println("data received:", string(data))
			})
			return err
		},
	}
}

// TestCmd defines the Xojo project `test` command.
func TestCmd() cli.Cmd {
	return cli.Cmd{
		Name:    "test",
		Summary: "Runs Xojo project tests after a project has been started via `run` command.",
		Handler: func(ctx *cli.CmdContext) error {
			// TODO: implement Xojo test's requests and responses.
			return nil
		},
	}
}
