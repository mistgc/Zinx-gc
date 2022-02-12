package znet

import (
	"Zinx/ziface"
	"fmt"
	"net"
	"Zinx/utils"
)

// iServer的接口实现，定义一个Server的服务器模块
type Server struct {
	// Name of server
	Name string

	// Ip version of server
	IPVersion string

	// Ip of the server listening
	IP string

	// Port of the server listening
	Port int

	// Add a router to current Server
	Router ziface.IRouter
}

// 服务器启动
func (s *Server) Start() {
	fmt.Printf("[Zinx] Server name: %s, listenner at IP: %s, Port: %d is starting...\n",
		utils.GlobalObject.Name, 
		utils.GlobalObject.Host,
		utils.GlobalObject.TcpPort,
	)
	fmt.Printf("[Zinx] Version %s, MaxConn: %d, MaxPackageSize: %d",
		utils.GlobalObject.Version,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPackageSize,
	)

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
			dealConn := NewConnection(conn, cid, s.Router)
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

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("Add Router Succ!")
}

// 初始化Server模块
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		Router:	   nil,
	}

	return s
}
