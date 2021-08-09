package main

import (
	"context"
	"fmt"
	"golang-grpc/average/protoaverage"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("I'm Client")

	conn, err := grpc.Dial("localhost:8010", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := protoaverage.NewAverageServiceClient(conn)

	// doUnary(c)

	// doStreaming(c)

	doClientStreaming(c)
}

func doClientStreaming(c protoaverage.AverageServiceClient) {

	req := []*protoaverage.AverageRequest{
		&protoaverage.AverageRequest{
			Average: &protoaverage.Average{
				A: 1,
			},
		},
		&protoaverage.AverageRequest{
			Average: &protoaverage.Average{
				A: 2,
			},
		},
		&protoaverage.AverageRequest{
			Average: &protoaverage.Average{
				A: 3,
			},
		},
		&protoaverage.AverageRequest{
			Average: &protoaverage.Average{
				A: 4,
			},
		},
	}

	stream, err := c.Average(context.Background())

	if err != nil {
		panic(err)
	}

	for _, request := range req {
		stream.Send(request)
		fmt.Println(request)
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Result)

}
