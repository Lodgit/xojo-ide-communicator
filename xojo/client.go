package xojo

import (
	"fmt"

	"github.com/joseluisq/gonetc"
)

// Client defines the Xojo IDE client communicator.
type Client struct {
	sock         *gonetc.NetClient
	protoVersion string
	ProjectCmds  *ProjectCommands
}

// New creates a new Xojo client instance.
func New(xojoSocketPath string) *Client {
	sock := gonetc.New("unix", xojoSocketPath)
	protoVersion := fmt.Sprintf("{\"protocol\":%d}%s", XojoCommunicationProtocolVersion, XojoNullChar)
	return &Client{
		sock:         sock,
		protoVersion: protoVersion,
		ProjectCmds: &ProjectCommands{
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
