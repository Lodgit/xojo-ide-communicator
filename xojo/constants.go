// +build linux

package xojo

// XojoUnixSocketPath specifies Xojo unix socket path on Linux systems
const XojoUnixSocketPath = "/tmp/XojoIDE"

// XojoCommunicationProtocolVersion specifies Xojo IDE communcation protocol
const XojoCommunicationProtocolVersion = 2

// XojoTestingServerAddress specifies Xojo Testing TCP server address and port
const XojoTestingServerAddress string = "127.0.0.1:8123"

// XojoNullChar specifies an ASCII NULL char (null terminator) in order to signal a transmission's end.
const XojoNullChar string = "\x00"
