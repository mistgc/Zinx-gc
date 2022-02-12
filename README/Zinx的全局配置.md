# Zinx的全局配置

随着架构越来越大，参数越来越多，为了省去我们后续高频率地修改参数，我们需要实现一个用于加载配置的模块，和一个全局获取Zinx参数的对象。

## 配置加载模块

以 `json` 格式：

`conf/zinx.json` （由用户填写）

```json
{
    "Name":"demo server",
    "Host":"127.0.0.1",
    "TcpPort":"7777",
    "MaxConn":3
}
```

#### 1.创建全局参数文件

创建 `zinx/utils` 文件夹，在下面创建 `globalobj.go` 文件。  

#### 2.提供init初始化方法

通过提供的 `init()` 方法，初始化 `GlobalObject` 对象，和加载服务端应用配置文件 `conf/zinx.json`

#### 3.硬参数替换与Server初始化参数配置

