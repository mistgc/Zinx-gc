# Zinx 基础路由模块 (router module)

- Request请求封装

  > 将连接与数据保存在一起

  - 属性
    - 连接 `conn ziface.IConnection`
    - 请求数据 `data []byte`

  - 方法

    - 获取当前连接 `GetConnection() ziface.IConnection`

    - 获得当前数据 `GetData() []byte`

      

- Router模块
  - 抽象的Router
    - 处理业务之前的方法 `PreHandle(request IRequset)`
    - 处理业务的主方法 `Handle(request IRequset)`
    - 处理业务之后的方法 `PostHandle(request IRequset)`
  - 具体的BaseRouter
    - 处理业务之前的方法 `PreHandle(request ziface.IRequset)`
    - 处理业务的主方法 `Handle(request ziface.IRequset)`
    - 处理业务之后的方法 `PostHandle(request ziface.IRequset)`

## Zinx 集成router模块

- IServer增添路由添加方法
  - `AddRouter(router IRouter)`
- Server类增添Router成员

- Connection类绑定一个Router成员

- 在Connection调用已经注册的Router处理业务
