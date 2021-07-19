#### 1、整体框架

服务端 :基于Gorilla/Websocket协议的在线聊天服务，通信协议使用Protobuf
客户端：使用fyne开发聊天客户端

#### 2、目录结构

```
.
├── README.md
├── __pycache__
│   └── locustfile.cpython-39.pyc
├── app
│   └── mian.go
├── go.mod
├── go.sum
├── internal
│   ├── ctrl
│   │   └── chatCtrl.go
│   ├── log
│   │   ├── logConfig.go
│   │   └── sys.log
│   ├── model
│   │   ├── ChatProto.pb.go
│   │   ├── ChatProto.proto
│   │   └── constantInfo.go
│   ├── view
│   │   └── chatView.go
│   └── wsClient
│       └── chatClient.go
└── 客户端流程图.png


```

#### 3. 代码逻辑分层


|层|文件夹|主要职责|调用关系|其他说明|
| ------------ | ------------ | ------------ | ------------ | ------------ |
|应用层 |app/main.go  |服务器启动 |调用view层   |不可同层调用
|ctrl层  |internal/ctrl | 处理来具体业务逻辑| 调用websocket层，被view调用  |不可同层调用
|websocket层 |internal/wsClient|提供基础的websocket功能 | 调用model，被ctrl层调用    |不可同层调用
| model |internal/model  |定义数据类型 | 被websocket层   |不可同层调用
| 配置文件 |internal/config  |日志配置 | 被websocket层 service层调用   |不可同层调用

#### 4.存储设计

```
message ChatRequest {
    string userName = 1;
    string type = 2;
    string content = 3;
    map<string,string> userList = 4;
}
```



#### 5. 第三方库

1. websocket框架
   代码 https://github.com/gorilla/websocket
   文档 https://pkg.go.dev/github.com/gorilla/websocket

2. 用于 数据传输
   代码 https://github.com/protocolbuffers/protobuf
   文档 https://developers.google.com/protocol-buffers/docs/gotutorial

3. 用于 客户端界面构建
   代码 https://github.com/fyne-io/fyne
   文档 https://developer.fyne.io/

#### 6. 如何编译执行

go run main.go

#### 7.todo

页面过于简单，数据与页面交互部分优化













