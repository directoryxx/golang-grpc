package main

import (
	"context"
	"fmt"
	"golang-grpc/squareroot/protosquareroot"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("I'm Client")

	conn, err := grpc.Dial("localhost:8010", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := protosquareroot.NewPrimeServiceClient(conn)

	// Correct Request
	// req := &protosquareroot.SquareRootRequest{
	// 	Number: 10,
	// }

	// Incorrect request
	req := &protosquareroot.SquareRootRequest{
		Number: -1,
	}

	res, err := c.SquareRoot(context.Background(), req)
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("Make sure input didnt negative number")
			}
		}
	} else {
		fmt.Println(res.Result)
	}

}
