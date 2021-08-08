package main

import (
	"context"
	"fmt"
	"golang-grpc/sum/protosum"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("I'm Client")

	conn, err := grpc.Dial("localhost:8010", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := protosum.NewSumServiceClient(conn)

	doUnary(c)
}

func doUnary(c protosum.SumServiceClient) {
	req := &protosum.SumRequest{
		Sum: &protosum.Sum{
			A: 2,
			B: 3,
		},
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Result)
}
