package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"time"
	"xojoidecom/xojo"
	"xojoidecom/xojotesting"

	cli "github.com/joseluisq/cline"
	"github.com/joseluisq/gonetc"
)

// RunCmd defines the Xojo project `run` command.
func RunCmd() cli.Cmd {
	return cli.Cmd{
		Name:    "run",
		Summary: "Runs a Xojo project in debug mode. Example: xojo-ide-com run [OPTIONS] PROJECT_FILE_PATH",
		Flags: []cli.Flag{
			cli.FlagInt{
				Name:    "delay",
				Summary: "Workaround delay in seconds to wait until the current application is finally displayed on screen.",
				Value:   5,
				Aliases: []string{"d"},
			},
			cli.FlagBool{
				Name:    "with-tests",
				Summary: "Run available unit tests through XojoUnit testing server. The testing TCP server should run on " + xojo.XojoTestingServerAddress,
				Aliases: []string{"t"},
			},
		},
		Handler: func(ctx *cli.CmdContext) error {
			runTests := ctx.Flags.Bool("with-tests").IsProvided()
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
			xo := xojo.New(xojo.XojoUnixSocketPath)
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
			// 5. Run tests if option is available
			if runTests {
				// 6. Xojo TCP connection
				log.Println("running project tests via XojoUnit...")
				client := gonetc.New("tcp", xojo.XojoTestingServerAddress)
				if err := client.Connect(); err != nil {
					log.Println("project XojoUnit Testing server is not available")
					return err
				}
				var testResultb []byte
				client.Listen(func(data []byte, err error, done func()) {
					if err != nil {
						log.Println(err)
						log.Fatalln(string(data))
					}
					testResultb = append(testResultb, data...)
					if len(data) == 0 || bytes.HasSuffix(data, []byte(xojo.XojoNullChar)) {
						done()
					}
				})
				// 6.1 Parse JSON test result
				testResultb = bytes.TrimSuffix(testResultb, []byte(xojo.XojoNullChar))
				testResult, err := xojotesting.ParseTestResult(testResultb)
				if err != nil {
					log.Fatalln(err)
				}
				// 6.2 Display test result in readable format
				fmt.Println()
				fmt.Printf("     XojoUnit version: %s\n", testResult.XojoUnitVersion)
				fmt.Printf("     Xojo IDE version: %s\n", testResult.XojoVersion)
				fmt.Printf("           Start time: %s\n", testResult.StartTime)
				fmt.Printf("          Total tests: %s\n", testResult.Total)
				fmt.Printf("         Failed tests: %s\n", testResult.Failures)
				fmt.Printf("         Passed tests: %s\n", testResult.PassedCount)
				fmt.Printf("        Skipped tests: %s\n", testResult.Skipped)
				fmt.Printf("        Skipped tests: %s\n", testResult.NotImplemented)
				fmt.Printf("Not implemented tests: %s\n", testResult.NotImplemented)
				fmt.Println()
				var hasTestsFailed bool = false
				for _, g := range testResult.Groups {
					fmt.Printf("=== RUN   %s\n", g.Name)
					for _, t := range g.Tests {
						if t.Passed {
							fmt.Printf("--- PASS: %s/%s (%s)\n", g.Name, t.Name, t.Duration)
						} else {
							fmt.Printf("--- FAIL: %s/%s (%s)\n", g.Name, t.Name, t.Duration)
						}
					}
					if g.Failures == 0 {
						fmt.Printf("PASS: %s (%s)\n", g.Name, g.Duration)
					} else {
						hasTestsFailed = true
						fmt.Printf("FAIL: %s (%s)\n", g.Name, g.Duration)
					}
					fmt.Printf("Total tests: %d\n", g.Total)
					fmt.Printf("Not implemented: %d\n", g.NotImplemented)
					fmt.Printf("Passed tests: %d\n", g.PassedCount)
					fmt.Printf("Failed tests: %d\n", g.Failures)
					fmt.Printf("Skipped tests: %d\n", g.Skipped)
					fmt.Println()
				}
				if hasTestsFailed {
					fmt.Printf("✗ Tests has been failed.\n")
					fmt.Println()
					os.Exit(1)
				} else {
					fmt.Println("✓ All tests passed successfully!")
					fmt.Println()
				}
				// 7. Close current project
				err = xo.ProjectCmds.Close(func(data []byte, err error) {
					if err != nil {
						log.Println(err)
						log.Fatalln(string(data))
					}
					log.Println("data received:", string(data))
				})
			}
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
			xo := xojo.New(xojo.XojoUnixSocketPath)
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
