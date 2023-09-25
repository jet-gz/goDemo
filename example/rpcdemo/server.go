package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
)

func hellow(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hellow world")
}

type Panda int

func (this *Panda) Getinfo(argType int, replyType *int) error {

	fmt.Println("打印", argType)
	*replyType = argType + 10086
	return nil
}

func main() {

	http.HandleFunc("/hellow", hellow)

	pd := new(Panda)
	// 注册
	rpc.Register(pd)
	rpc.HandleHTTP()

	ln, err := net.Listen("tcp", ":10086")
	if err != nil {
		fmt.Println("网络错误")
	}
	http.Serve(ln, nil)

}
