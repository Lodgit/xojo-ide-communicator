package xojotesting

import (
	"bytes"
	"xojoidecom/xojo"

	"github.com/joseluisq/gonetc"
)

// Client defines XojoUnit Testing client.
type Client struct {
	sock *gonetc.NetClient
}

// New creates a new Xojo Testing client instance.
func New(xojoSocketPath string) *Client {
	sock := gonetc.New("tcp", xojoSocketPath)
	return &Client{
		sock: sock,
	}
}

// Connect tries to connect to Xojo Xojo Unit Testing server and triggers the tests immediately.
func (c *Client) Connect() (err error) {
	return c.sock.Connect()
}

// Listen listens for a single JSON response with tests result.
func (c *Client) Listen() (data []byte, err error) {
	c.sock.Listen(func(xdata []byte, xerr error, done func()) {
		if xerr != nil {
			err = xerr
			return
		}

		data = append(data, xdata...)
		if len(xdata) == 0 || bytes.HasSuffix(xdata, []byte(xojo.XojoNullChar)) {
			data = bytes.TrimSuffix(data, []byte(xojo.XojoNullChar))
			done()
		}
	})
	return
}

// Close closes the current Xojo IDE socket.
func (c *Client) Close() {
	c.sock.Close()
}
