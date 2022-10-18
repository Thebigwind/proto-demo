package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"lxf.com/proto-demo/service"
)

func main() {
	user := &service.User{
		Username: "lxf",
		Age:      20,
	}
	data, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}
	//反序列号
	newUser := &service.User{}
	err = proto.Unmarshal(data, newUser)
	if err != nil {
		panic(err)
	}
	fmt.Printf("newUser:%v", newUser.String())
}
