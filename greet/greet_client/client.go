package main

import (
	"fmt"
	"golang-grpc/greet/greetpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("I'm Client")

	conn, err := grpc.Dial("localhost:8010", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	// fmt.Println("Client Created : ", &c)
	fmt.Printf("Client Created : %f", c)
}
