package xojo

import (
	"log"

	"github.com/joseluisq/goipcc"
)

// Commands defines a command for run a Xojo project.
type Commands struct {
	sock *goipcc.IPCSockClient
}

// Run runs the current opened project.
func (c *Commands) Run(handler func(data []byte, err error)) error {
	str := "{\"tag\":\"build\",\"script\":\"DoCommand(\\\"RunApp\\\")\nprint \\\"App is running.\\\"\"}\x00"
	log.Println("run project command sent:", str)
	_, err := c.sock.Write([]byte(str), handler)
	if err != nil {
		return err
	}
	return nil
}

// Close closes the current opened project.
func (c *Commands) Close(handler func(data []byte, err error)) error {
	str := "{\"tag\":\"build\",\"script\":\"CloseProject(False)\nprint \\\"App is closed.\\\"\"}\x00"
	log.Println("close project sent:", str)
	_, err := c.sock.Write([]byte(str), handler)
	if err != nil {
		return err
	}
	return nil
}

// Build builds current opened project.
func (c *Commands) Build(handler func(data []byte, err error)) error {
	// TODO: customize `BuildApp` args
	str := "{\"script\":\"Print BuildApp(16,False)\", \"tag\":\"build\"}\x00"
	log.Println("build project command sent:", str)
	_, err := c.sock.Write([]byte(str), handler)
	if err != nil {
		return err
	}
	return nil
}
