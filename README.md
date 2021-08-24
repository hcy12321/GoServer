# GoServer

简易的基于golang实现的游戏服务器。

## 模块简要说明
----

已实现功能模块如下：
1. tcp服务封装
2. protobuf解析及注册
3. 路由(手动注册版本)
4. 供测试使用的tcp client(client.go默认注释)
5. 测试逻辑(logic文件夹)
6. 简易Dockerfile

游戏服务器理应包含还未实现的其他功能：
1. 缓存数据库模块(如Redis)
2. 持久化数据库模块(如mongodb)
3. 配置表读写工具模块
4. 工具函数模块
5. 其他网络协议实现(udp,ws,http等等)

后期应实现的模块
1. 压测工具
2. GM工具

## 设计概述
----

依赖容器，单线程，多节点进程，无状态


## 数据简易说明
----

传输的每个包数据，前4个字节为后面包长度，5-8字节表示协议id，然后是包体数据

## 使用
----

* go env -w GO111MODULE=on
* go env -w GOPROXY=https://goproxy.cn,direct
* go build main.go
* 运行生成的可执行文件

如果要运行测试的client程序：
* go build client.go
* 运行生成的可执行程序

## protobuf修改及使用
----
* 所有proto文件都应在proto目录下
* 确认已下载protoc执行文件并加入环境变量
* 修改协议结构应遵循首个参数为BaseMessage类型并且命名为base
* 修改完成后执行目录下gen
* 在main中调用convert.RegisterProto注册处理结构


## 增加协议处理函数
----
* 在logic目录下增加文件或者选择一个已有文件
* 在文件内实现一个函数，需遵循MessageHandler格式
* 在main中调用router的registerHandler来注册调用