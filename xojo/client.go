package xojo

import (
	"fmt"

	"github.com/joseluisq/goipcc"
)

// Client defines the Xojo IDE client communicator.
type Client struct {
	sock         *goipcc.IPCSockClient
	program      *program
	protoVersion string
	Commands     *Commands
}

// New creates a new Xojo client instance.
func New() *Client {
	sock := goipcc.New(xojoUnixSocketPath)
	protoVersion := fmt.Sprintf("{\"protocol\":%d}\x00", xojoCommunicationProtocolVersion)
	return &Client{
		sock:         sock,
		protoVersion: protoVersion,
		program: &program{
			execPath: xojoExecFile,
		},
		Commands: &Commands{
			sock: sock,
		},
	}
}

// Connect tries to connect to Xojo IDE socket and sets the Xojo communication protocol.
func (c *Client) Connect() error {
	err := c.sock.Connect()
	if err != nil {
		return err
	}
	_, err = c.sock.Write([]byte(c.protoVersion), nil)
	if err != nil {
		return err
	}
	return nil
}

// Close closes the current Xojo IDE socket.
func (c *Client) Close() {
	c.sock.Close()
}
