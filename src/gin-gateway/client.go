package main

import (
	"context"
	"fmt"

	proto "gin-gateway/routes"

	micro "github.com/micro/go-micro"
)

func main() {
	// Create a new service
	service := micro.NewService(micro.Name("greeter.client"))
	// Initialise the client and parse command line flags
	service.Init()

	// Create new greeter client
	greeter := proto.NewGreeterService("greeter", service.Client())

	// Call the greeter
	str := "john"
	rsp, err := greeter.Hello(context.TODO(), &proto.Request{Name: &str})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(*rsp.Msg)
}
