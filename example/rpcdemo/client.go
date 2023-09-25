package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 建立网络连接
	cli, err := rpc.DialHTTP("tcp", "127.0.0.1:10086")
	if err != nil {
		fmt.Println("网络错误")
	}
	var pd int
	err = cli.Call("Panda.Getinfo", 10086, &pd)
	if err != nil {
		fmt.Println("错误")
	}

	fmt.Println("得到的值", pd)

}
