package main

import (
	"context"
	"fmt"
	"golang-grpc/findmax/protofindmax"
	"io"
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

	c := protofindmax.NewFindmaxServiceClient(conn)

	waitc := make(chan struct{})

	req, err := c.Findmax(context.Background())

	if err != nil {
		panic(err)
	}

	go func() {
		numbers := []int{4, 7, 2, 19, 4, 6, 32}
		for _, num := range numbers {
			req.Send(&protofindmax.FindmaxRequest{
				Findmax: &protofindmax.Findmax{
					A: int64(num),
				},
			})
			time.Sleep(100 * time.Microsecond)
		}
		req.CloseSend()
	}()

	go func() {
		for {
			res, err := req.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
				// break
			}
			maximum := res.GetResult()
			fmt.Println(maximum)
		}
		close(waitc)
	}()

	<-waitc
}
