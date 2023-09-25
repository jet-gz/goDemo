package main

import (
	"context"
	"fmt"
	"grpcdemo/pb"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

// 相当于接受请求，客户端调用它    context 客户端传过来的， in 客户端传过来的参数
func (this *server) SayHellow(ctx context.Context, in *pb.HellowReq) (*pb.HellowRes, error) {

	fmt.Println("SayHellow收到消息：", in)
	return &pb.HellowRes{Msg: "你好" + in.Name}, nil
}
func (this *server) GetName(ctx context.Context, in *pb.NameReq) (*pb.NameRes, error) {
	fmt.Println("GetName收到消息：", in)
	return &pb.NameRes{Name: in.Name + "->Jet"}, nil
}

func main() {

	ln, err := net.Listen("tcp", ":6000")
	if err != nil {
		fmt.Println("网络错误")
	}
	// 创建grpc服务
	svr := grpc.NewServer()
	//注册服务
	pb.RegisterHellowServerServer(svr, &server{})
	// 等待网络连接
	svr.Serve(ln)
}
