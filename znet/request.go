package znet

import (
	"Zinx/ziface"
)

type Request struct {
	// The connection has been established with the client.
	conn ziface.IConnection

	// The data requested by the client.
	data []byte
}

// Get current connection
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// Get request data
func (r *Request) GetData() []byte {
	return r.data
}
