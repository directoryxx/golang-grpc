package main

import (
	"context"
	"fmt"
	"golang-grpc/greet/greetpb"
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

	c := greetpb.NewGreetServiceClient(conn)

	// doUnary(c)

	// doStreaming(c)

	doClientStreaming(c)
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

func doStreaming(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Angga",
			LastName:  "Wijaya",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
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

func doClientStreaming(c greetpb.GreetServiceClient) {

	req := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Angga1",
				LastName:  "W1",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Angga2",
				LastName:  "W2",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Angga3",
				LastName:  "W3",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Angga4",
				LastName:  "W4",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())

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
