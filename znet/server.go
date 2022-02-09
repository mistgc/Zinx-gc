package znet

import (
	"Zinx/ziface"
	"fmt"
	"net"
	"errors"
)

// iServer的接口实现，定义一个Server的服务器模块
type Server struct {
	// 服务器名称
	Name string

	// 服务器ip版本
	IPVersion string

	// 服务器监听的ip
	IP string

	// 服务器监听的端口
	Port int

}

func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[Conn Handle] CallBackToClient...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("Write back buffer error ", err)
		return errors.New("CallBackToClient error")
	}

	return nil
}

// 服务器启动
func (s *Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP :%s, Port :%d, is starting\n", s.IP, s.Port)

	go func(){
		// 1.Get an Addr of TCP.
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("Resolve tcp addr error:", err)
			return
		}

		// 2.Listen address of server.
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("Listen ", s.IPVersion, "err", err)
		}

		fmt.Println("Start Zinx server successfully, ", s.Name, "successful, Linstenning...")
		var cid uint32
		cid = 0

		// 3.Block to wait for client connections and handle client connection business. (I/O)
		for {
			// If have connected by a client, blocking is returned.
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}

			// If the connection is successful, the business begins.
			// Register business 'CallBackToClient'.
			dealConn := NewConnection(conn, cid, CallBackToClient)
			cid++

			// Start current connection business.
			go dealConn.Start()

		}
	}()
}

// 服务器停止
func (s *Server) Stop() {
	// TODO: Stop or recycle some server resources, status, or some established connection information.

}

// 服务器运行
func (s *Server) Serve() {
	s.Start()

	// TODO: Do other business after the server starts.

	// Blocking status
	select{}
}

// 初始化Server模块
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}

	return s
}
