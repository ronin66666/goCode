# gRPC 和 Protobuf

## gRPC

### RPC

RPC指远程过程调用（`Remote Procedure Call`）,它调用包含传输协议和编码（对象序列）协议等，允许运行于一台计算机上的程序调用另一台计算机上的程序，而开发人员无需额外为这个交互作用编程，就像对本地函数进行调用一样方便

### gRPC

gRPC是一个高新能、开源、通用的RPC框架，基于HTTP/2标准设计，拥有双向流、流控、头部压缩、单TCP连接上的多复用请求特性，这些特性使得其在移动设备上表现更好、更节省空间。

gRPC的的接口描述语言使用的是Protobuf是由Google开源的

## Protobuf

是一种与语言、平台无关、且可扩展的序列化结构化数据的数据描述语言，通常称其为IDL，常用于通信协议、数据储存等，与JSON、XML相比，它更小、更快。

### protobuf使用

文档：https://developers.google.com/protocol-buffers/docs/reference/go-generated

需要先安装protoc编译器，主要编译.proto文件

检查是否安装成功

```bash
protoc --version
```

protoc插件

针对不同的语言，还需安装运行时的protoc插件，go语言的是protoc-gen-go插件

```bash
go get -u github.com/golang/protobuf/protoc-gen-go@v1.3.2
```

将所编译安装的protoc插件的可执行文件移到相应的bin目录下

```bash
mv $GOPATH/bin/protoc-gen-go /usr/local/bin
```

这个命令不是必须得，其主要的目的是将二进制文件protoc-gen-go移到bin目录下，让其可以直接运行progoc-gen-to插件，也可以直接配置环境变量

## 编译和生成proto文件

### 创建proto文件

protobuf支持数据类型有：double、float、int32、int64、uint32、uint64、sint32、sint64、fixed32、fixed64、sfixed32、sfixed64、bool、string、bytes

```protobuf
syntax = "proto3";
package helloworld;

service Greeter {
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

message HelloRequest {
  //类型 字段名 字段编号
  string name = 1;
}

message HelloReply {
  string message = 1;
}
```

### 生成proto文件

项目根目录下，执行下面命令，会在.proto同目录中生成.pb.go代码，需要注意的是**.proto中的包名要跟生成的go文件所在的包名相同**，不然会报错

```bash
$ protoc --go_out=plugins=grpc,paths=source_relative:. ./proto/helloworld.proto 
```

或`./proto/*.proto`将目录下所有的.proto文件生成.pb.go文件

```
protoc --go_out=plugins=grpc,paths=source_relative:. ./proto/*.proto 
```



























