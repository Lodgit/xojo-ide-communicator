package xojo

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"sync"
	"testing"
	"time"
)

// unixSocketPath is a default Unix socket path used on tests.
const unixSocketPath = "/tmp/mysocket"

// unixSocketDelay defines milliseconds pause in order to wait until
// the listening server (`socat`) is ready to accept connections.
const unixSocketDelay = 500

// listeningSocket defines a listening unix socket.
type listeningSocket struct {
	cmd *exec.Cmd
	wg  *sync.WaitGroup
}

// createListeningSocket creates a new listening unix socket using `socat` tool.
func createListeningSocket() (*listeningSocket, error) {
	exec.Command("rm", "-rf", unixSocketPath).Run()

	var out bytes.Buffer
	cmd := exec.Command("socat", "UNIX-LISTEN:"+unixSocketPath+",fork", "exec:'/bin/cat'")
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = &out
	if err := cmd.Start(); err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go cmd.Wait()
	time.Sleep(unixSocketDelay * time.Millisecond)

	return &listeningSocket{
		wg:  &wg,
		cmd: cmd,
	}, nil
}

// close method closes current socket connection signaling it to finish.
func (s *listeningSocket) close() error {
	return s.cmd.Process.Signal(os.Interrupt)
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Client
	}{
		{
			name: "valid instance",
			want: &Client{
				protoVersion: fmt.Sprintf("{\"protocol\":%d}\x00", XojoCommunicationProtocolVersion),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(unixSocketPath); tt.want.protoVersion != got.protoVersion || got.sock == nil || got.ProjectCmds == nil {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Connect(t *testing.T) {
	lsock, err := createListeningSocket()
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	tests := []struct {
		name       string
		unixSocket string
		wantErr    bool
	}{
		{
			name:    "invalid socket connection",
			wantErr: true,
		},
		{
			name:       "valid socket connection",
			unixSocket: unixSocketPath,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(tt.unixSocket)
			if err := c.Connect(); (err != nil) != tt.wantErr {
				t.Errorf("Client.Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if c.sock == nil {
					t.Errorf("Client.Connect() = sock: %v, want not nil", c.sock)
				}
				if c.ProjectCmds == nil {
					t.Errorf("Client.Connect() = ProjectCmds: %v, want not nil", c.ProjectCmds)
				}
			}
		})
	}
	if err := lsock.close(); err != nil {
		t.Errorf("%v", err)
		return
	}
}

func TestClient_Close(t *testing.T) {
	lsock, err := createListeningSocket()
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	tests := []struct {
		name       string
		unixSocket string
	}{
		{
			name:       "close a valid socket connection",
			unixSocket: unixSocketPath,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(tt.unixSocket)
			if err := c.Connect(); err != nil {
				t.Errorf("Client.Connect() error = %v", err)
			}
			c.Close()
		})
	}
	if err := lsock.close(); err != nil {
		t.Errorf("%v", err)
		return
	}
}
