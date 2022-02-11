package ziface

import (
	"net"
)

// Defines the abstraction layer of the connection module.
type IConnection interface {
	// Start connection
	Start()

	// Stop  connection
	Stop()

	// Get the socket to which the current connection is bound.
	GetTCPConnection() *net.TCPConn

	// Get the ID of the current connection module.
	GetConnID() uint32

	// Get TCP stutas of the remote client. (IP port)
	RemoteAddr() net.Addr

	// Send data to the client.
	Send(data []byte) error
}

// Defines a method to handle the connection business.
type HandleFunc func(*net.TCPConn, []byte, int) error
