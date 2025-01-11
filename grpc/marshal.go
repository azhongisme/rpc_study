package main

import (
	"fmt"
	"grpc_study/service"

	"google.golang.org/protobuf/proto"
)

func main() {
	user := &service.User{
		Username: "azhong",
		Age: 18,
	}

	marshal, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}

	newUser := &service.User{}
	err = proto.Unmarshal(marshal, newUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(newUser.String())
}