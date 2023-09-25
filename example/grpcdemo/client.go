package main

import (
	"context"
	"fmt"
	"grpcdemo/pb"

	"google.golang.org/grpc"
)

func main() {
	// 建立网络连接
	conn, err := grpc.Dial("127.0.0.1:6000", grpc.WithInsecure())
	if err != nil {
		fmt.Println("网络错误")
	}

	// 过去grpc句柄
	g := pb.NewHellowServerClient(conn)

	r1, _ := g.SayHellow(context.Background(), &pb.HellowReq{Name: "Jet "})
	fmt.Println(r1)

	r2, _ := g.GetName(context.Background(), &pb.NameReq{Name: "ggg"})
	fmt.Println(r2)

}
