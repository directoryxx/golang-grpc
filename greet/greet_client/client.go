package main

import (
	"context"
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

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetingRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Angga",
			LastName:  "Wijaya",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Result)
}
