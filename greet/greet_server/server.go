package main

import (
	"fmt"
	"golang-grpc/greet/greetpb"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Printf("Hello World")

	lis, err := net.Listen("tcp", "0.0.0.0:8010")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
