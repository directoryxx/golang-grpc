package main

import (
	"context"
	"fmt"
	"golang-grpc/prime/protoprime"
	"io"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("I'm Client")

	conn, err := grpc.Dial("localhost:8010", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := protoprime.NewPrimeServiceClient(conn)
	req := &protoprime.PrimeRequest{
		Prime: &protoprime.Prime{
			A: 120,
		},
	}

	resStream, err := c.Prime(context.Background(), req)
	if err != nil {
		panic(err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// Catch End of file
			break
		}

		if err != nil {
			panic(err)
		}

		fmt.Println(msg.Result)
	}
}
