package znet

import (
	"fmt"
	"net"
	"Zinx/ziface"
)

// Connection module
type Connection struct {
	// TCP socket of current connection.
	Conn *net.TCPConn

	// Connection ID
	ConnID uint32

	// Current connection stutas
	isClosed bool

	// Mathod of processing business to which the current connection is bound.
	handleAPI ziface.HandleFunc

	// Channel that tells the current connection that it has exited or stopped.
	ExitChan chan bool
}

// Initialize the mathod of connection module
func NewConnection(conn *net.TCPConn, connID uint32, callbackAPI ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn: conn,
		ConnID: connID,
		handleAPI: callbackAPI,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}

	return c
}

// Read data serve for connection
func (c *Connection) StartReader(){
	fmt.Println("Reader Goroutine is running...")
	
	defer fmt.Println("ConnID = ", c.ConnID, " Reader is exit, remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		// Read the data from client to the buffer. MaxSize = 512 Byte
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("Receive buffer error ", err)
			continue
		}

		// Call the 'HandleAPI' to which the current connection is bound.
		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("ConnID = ", c.ConnID, " handle is error ", err)
			break
		}
	}
}

// TODO: Write data serve for connection

// Start connection
func (c *Connection) Start() {
	fmt.Println("Connection Start... ConnID = ", c.ConnID)

	// Start the read data serve for current connection.
	go c.StartReader()

	// TODO: Start the write data serve for current connection.
}

// Stop  connection
func (c *Connection) Stop() {
	fmt.Println("Connection Stop... ConnID = ", c.ConnID)

	if c.isClosed == true {
		return
	}

	c.isClosed = true

	// Close socket connection
	c.Conn.Close()

	// Recycle
	close(c.ExitChan)
}

// Get the socket to which the current connection is bound.
func (c *Connection) GetConnection() *net.TCPConn {
	return c.Conn
}

// Get the ID of the current connection module.
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// Get TCP stutas of the remote client. (IP port)
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// Send data to the client.
// func (c *Connection) Send(data []byte) error {
// 	
// }
