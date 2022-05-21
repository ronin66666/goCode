package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc/proto"
	"io"
)
var port string

func init() {
	port = "8000"
}
func main() {
	conn, err := grpc.Dial(":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	client := proto.NewGreeterClient(conn)
	//err = SayHello(client)
	//err = SayList(client, &proto.HelloRequest{Name: "zhangsan"})
	//err = SayRecord(client, &proto.HelloRequest{Name:"SayRecord"})
	err = SayRoute(client, &proto.HelloRequest{Name:"SayRoute"})
	if err != nil {
		fmt.Println(err)
	}
}

func SayHello(client proto.GreeterClient) error  {
	resp, err := client.SayHello(context.Background(), &proto.HelloRequest{Name: "zhangsan"})
	if err != nil {
		return err
	}
	fmt.Printf("client.SayHello resp: %s", resp.Message)
	return nil
}

//服务端流式传输
func SayList(client proto.GreeterClient, r *proto.HelloRequest)  error {
	stream, err := client.SayList(context.Background(), r)
	if err != nil {
		return err
	}
	for {
		resp, err := stream.Recv() //默认的MaxReceiveMessageSize值为1024 * 1024 * 4
		if err == io.EOF { //传输完毕后退出
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(resp.Message)
	}
	return nil
	//hello.list 0
	//hello.list 1
	//hello.list 2
	//hello.list 3
	//hello.list 4
	//hello.list 5
	//hello.list 6
}

//客户端流
func SayRecord(client proto.GreeterClient, r *proto.HelloRequest) error  {
	stream, err := client.SayRecord(context.Background())
	if err != nil {
		return err
	}
	for n := 0; n < 6; n++ {
		if err := stream.Send(r); err != nil {
			return err
		}
	}
	//接收发送，接收数据
	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	fmt.Printf("resp = %v \n", resp)
	return nil
}

//双向流
func SayRoute(client proto.GreeterClient, r *proto.HelloRequest) error  {
	stream, err := client.SayRoute(context.Background())
	if err != nil {
		return err
	}
	for n := 0; n < 6; n                                                                                                                                                                                                                                                                                       ++ {
		r.Name = fmt.Sprintf("sayRoute %d", n)
		//发送数据
		if err := stream.Send(r); err != nil {
			return err
		}

		//接送数据
		resp, err := stream.Recv()
		fmt.Printf("resp %v \n", resp)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	//断开发送
	if err := stream.CloseSend(); err != nil {
		return err
	}
	return nil
}

