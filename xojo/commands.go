package xojo

import (
	"github.com/joseluisq/goipcc"
)

// runCommand defines a command for run a Xojo project.
type runCommand struct {
	sock *goipcc.IPCSockClient
}

func (c *runCommand) Exec(handler func(data []byte, err error)) error {
	_, err := c.sock.Write([]byte("{\"script\":\"DoCommand(\"RunApp\")\", \"tag\":\"run\"}\x00"), handler)
	if err != nil {
		return err
	}
	return nil
}

// buildCommand defines a command for build a Xojo project.
type buildCommand struct {
	sock *goipcc.IPCSockClient
}

func (c *buildCommand) Exec(handler func(data []byte, err error)) error {
	// TODO: customize `BuildApp` args
	_, err := c.sock.Write([]byte("{\"script\":\"Print BuildApp(16,False)\", \"tag\":\"build\"}\x00"), handler)
	if err != nil {
		return err
	}
	return nil
}
