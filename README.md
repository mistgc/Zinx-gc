# Zinx-轻量级TCP服务器框架 (学习ing)
该项目是作者对Zinx框架的学习重构。

## 模块 (Module)
* Zinx框架
	* 基础server模块
	* 基础router模块
	* 全局配置
	* 消息封装
	* 多路由模式
	* 读写协程分离
	* 消息队列及多任务
	* 链接管理
	* 连接属性配置
* 基于Zinx的应用案例

## 基础server模块

### 方法
* 初始化Server
* 启动服务器
  * 创建addr
  * 创建listenner
  * 处理客户端请求
* 停止服务器
* 运行服务器
  * 调用启动服务器`Start()`方法之后阻塞
  * 在阻塞期间处理其它业务

### 属性
* name名称
* 监听的IP
* 监听的端口
