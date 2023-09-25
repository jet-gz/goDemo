package main

import (
	"fmt"
	"protoDemo/pb"

	"google.golang.org/protobuf/proto"
)

func main() {

	s := &pb.Student{
		Name:    "jet",
		Scores:  []int32{3, 4, 68},
		Male:    true,
		Subject: map[string]int32{"age": 29, "height": 178},
	}

	// 序列化
	data, err := proto.Marshal(s)
	if err != nil {
		fmt.Println("proto encode error: ", err)
		return
	}

	fmt.Println(data)

	newStudent := &pb.Student{}
	//反序列化
	err = proto.Unmarshal(data, newStudent)
	if err != nil {
		fmt.Println("proto decode error: ", err)
	}

	fmt.Println(newStudent)

}
