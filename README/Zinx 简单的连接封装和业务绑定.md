# Zinx 简单的连接封装和业务绑定

- 链接的模块

  - 方法
    - 启动连接 `Strat()`
    - 停止连接 `Stop()`
    - 获取当前连接的conn对象 (socket对象) `GetTCPConnection() *net.TCPConn`
    - 获取连接ID `GetConnID() uint32`
    - 得到客户端连接的地址和端口 `RemoteAddr() net.Addr`
    - 发送数据的方法 `Send(data []byte) error`
    - 连接所绑定的处理业务的函数类型 `type HandleFunc func(*net.TCPConn, []byte, int) error`
  - 属性
    - Socket TCP套接字 `Conn *net.TCPConn`
    - 连接ID `ConnID uint32`
    - 当前连接状态 (是否已经关闭) `isClosed bool`
    - 与当前连接所绑定的处理业务方法 `handleAPI ziface.HandleFunc`
    - 等待连接被动退出的channel `ExitChan chan bool`

  