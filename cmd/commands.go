package cmd

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"
	"xojoidecom/xojo"
	"xojoidecom/xojotesting"

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
				log.Fatalln("Xojo project file path was not provided.")
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
			log.Printf("Waiting %d second(s) for the application's start...\n", delay)
			time.Sleep(time.Duration(delay) * time.Second)
			// 5. Run tests if option is available
			if runTests {
				// 6. Xojo TCP connection with Testing server
				log.Println("Running project tests via XojoUnit...")
				client := xojotesting.New(xojo.XojoTestingServerAddress)
				if err := client.Connect(); err != nil {
					log.Println("project XojoUnit Testing server is not available")
					return err
				}
				// 6.1. Listen, receive one response only and parse the JSON test results
				testResultb, err := client.Listen()
				if err != nil {
					log.Println("XojoUnit Testing server communication error")
					return err
				}

				// 6.2. Check for incoming Xojo runtime errors on tests
				testResultb, err = xojotesting.CheckForGlobalErrorResponse(testResultb, nil)
				if err != nil {
					log.Println("XojoUnit Testing server responses with a Xojo runtime error")
					runtimeErrResult, err := xojotesting.ParseRuntimeErrorResult(testResultb)
					if err != nil {
						log.Println("Can not parse XojoUnit Testing server \"Runtime Error\" response")
						return err
					}
					runtimeErr := runtimeErrResult.RuntimeError
					fmt.Printf("   Error Type: %s\n", runtimeErr.ErrorType)
					fmt.Printf(" Error Number: %d\n", runtimeErr.ErrorNumber)
					fmt.Printf("Error Message: %s\n", runtimeErr.ErrorMessage)
					fmt.Printf(" Error Reason: %s\n", runtimeErr.ErrorReason)
					fmt.Printf("  Error Stack:\n")
					for i, s := range runtimeErr.ErrorStack {
						if s != "" {
							fmt.Printf("            %d: %s\n", i, s)
						}
					}
					fmt.Println()
					fmt.Printf("✗ XojoUnit Testing server has been failed.\n")
					fmt.Println()
					os.Exit(1)
				}

				testResult, err := xojotesting.ParseTestResult(testResultb)
				if err != nil {
					log.Println("Can not parse XojoUnit Testing server \"Test Results\" response")
					return err
				}
				// 6.3. Display test result in readable format
				fmt.Println()
				fmt.Println(" XojoUnit Testing Server")
				fmt.Println("=========================")
				fmt.Printf("XojoUnit version: %s\n", testResult.XojoUnitVersion)
				fmt.Printf("Xojo IDE version: %s\n", testResult.XojoVersion)
				fmt.Printf("      Start time: %s\n", testResult.StartTime)
				fmt.Println()
				var hasTestsFailed bool = false
				for _, g := range testResult.Groups {
					fmt.Printf("=== RUN   %s\n", g.Name)
					for _, t := range g.Tests {
						if t.Passed {
							fmt.Printf("--- PASS: %s/%s (%s)\n", g.Name, t.Name, t.Duration)
						} else if t.Type == "failed" {
							fmt.Printf("--- FAIL: %s/%s (%s)\n", g.Name, t.Name, t.Duration)
							fmt.Printf("             %s\n", strings.ReplaceAll(t.FailedMessage, "\n", "\n             "))

							// Check for a RuntimeError in current TestResult if it's present
							if t.RuntimeError.ErrorType != "" {
								fmt.Printf("   Error Type: %s\n", t.RuntimeError.ErrorType)
								fmt.Printf(" Error Number: %d\n", t.RuntimeError.ErrorNumber)
								fmt.Printf("Error Message: %s\n", t.RuntimeError.ErrorMessage)
								fmt.Printf(" Error Reason: %s\n", t.RuntimeError.ErrorReason)
								fmt.Printf("  Error Stack:\n")
								for i, s := range t.RuntimeError.ErrorStack {
									if s != "" {
										fmt.Printf("            %d: %s\n", i, s)
									}
								}
							}
						} else {
							fmt.Printf("--- SKIP: %s/%s (%s)\n", g.Name, t.Name, t.Duration)
						}
					}
					if g.FailuresCount == 0 {
						fmt.Printf("PASS: %s (%s)\n", g.Name, g.Duration)
					} else {
						hasTestsFailed = true
						fmt.Printf("FAIL: %s (%s)\n", g.Name, g.Duration)
					}
					fmt.Printf("Total tests: %d\n", g.Total)
					fmt.Printf("Not implemented: %d\n", g.NotImplementedCount)
					fmt.Printf("Passed tests: %d\n", g.PassedCount)
					fmt.Printf("Failed tests: %d\n", g.FailuresCount)
					fmt.Printf("Skipped tests: %d\n", g.SkippedCount)
					fmt.Println()
				}

				fmt.Println("Final test results:")
				fmt.Printf("          Total tests: %s\n", testResult.Total)
				fmt.Printf("         Failed tests: %s\n", testResult.FailuresCount)
				fmt.Printf("         Passed tests: %s\n", testResult.PassedCount)
				fmt.Printf("        Skipped tests: %s\n", testResult.SkippedCount)
				fmt.Printf("        Skipped tests: %s\n", testResult.NotImplementedCount)
				fmt.Printf("Not implemented tests: %s\n", testResult.NotImplementedCount)
				fmt.Println()

				if hasTestsFailed {
					fmt.Printf("✗ Tests has been failed / %s\n", testResult.FailuresCount)
					fmt.Println()
					os.Exit(1)
				} else {
					fmt.Println("✓ All tests passed successfully!")
					fmt.Println()
				}
				// 7. Close Testing server client and Xojo current project
				client.Close()
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
		Summary: "Builds a Xojo project to specific target(s). Example: xojo-ide-com build [OPTIONS] PROJECT_FILE_PATH",
		Flags: []cli.Flag{
			cli.FlagStringSlice{
				Name:    "targets",
				Summary: "Operating systems with their architectures. Coma-separated list with one or more target pairs in lower case. E.g linux-amd64,darwin-arm64,windows-386.",
				Aliases: []string{"t"},
			},
			cli.FlagBool{
				Name:    "reveal",
				Value:   false,
				Summary: "Open the built application directory using the operating system file manager available.",
				Aliases: []string{"r"},
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
			targets := ctx.Flags.StringSlice("targets").Value()
			if len(targets) == 0 {
				log.Fatalln("no build tagets specified. Use --targets option")
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
			for _, target := range targets {
				opts := xojo.BuildOptions{
					Target: target,
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
