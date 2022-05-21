package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/proto"
	"io"
	"net"
)

type GreeterServer struct {}

//一元请求：单次请求+响应
func (s *GreeterServer) SayHello(ctx context.Context,r *proto.HelloRequest) (*proto.HelloReply, error)  {
	return &proto.HelloReply{
		Message:              "Hello world" + r.Name,
	}, nil
}

//服务端流式RPC：客服端发起一次请求，服务端通过流式响应多次发送数据集，客户端Recv接收数据集
func (s *GreeterServer) SayList(r *proto.HelloRequest, stream  proto.Greeter_SayListServer) error  {
	for n := 0; n <= 6; n++ {
		//send方法执行过程
		//消息体（对象序列化）
		//压缩序列化后的消息体
		//为正在传输的消息体增加5字节的header（标志位）
		//判断压缩+序列化后的消息体总字节长度是否大于预设的maxsizeMessageSize(预设值为math.MaxInt32)，若超出则提示错误
		//写入给流的数据集
		err := stream.Send(&proto.HelloReply{Message: fmt.Sprintf("hello.list %d", n)})
		if err != nil{
			 return err
		}
	}
	return nil
}

//客户端流式RPC
func (s *GreeterServer) SayRecord(stream proto.Greeter_SayRecordServer) error {
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			//接收结束后返回结果
			message := &proto.HelloReply{Message: "say.record end"}
			return stream.SendAndClose(message)
		}
		if err != nil {
			return err
		}
		fmt.Printf("resp: %v \n", resp)
	}
}

//双向流
//首个请求一定是客户端发起的，但具体的交互方式（谁先谁后、一次发多少、响应多少，什么时候关闭）则由程序编写的方式来确定（可以结合协程）
//下面假设双向流式按顺序发送的
func (s *GreeterServer) SayRoute(stream proto.Greeter_SayRouteServer) error {
	n := 0
	var respStr string
	for {
		//从客户端接收
		resp, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		n++
		fmt.Printf("resp = %v \n", resp)

		//将接收到的数据发送到客户端
		respStr = resp.Name
		if err := stream.Send(&proto.HelloReply{Message:respStr});  err != nil {
			return err
		}
	}
}

func main() {
	server := grpc.NewServer()
	//注册服务，这样在接收请求时，即可通过内部的"服务发现"发现该服务端接口，并进行逻辑处理
	proto.RegisterGreeterServer(server, &GreeterServer{})
	lis, _ := net.Listen("tcp", ":8000")
	err :=  server.Serve(lis)
	if err != nil {
		fmt.Println(err)
	}
}
