package main

import (
	"context"
	"fmt"
	"golang-grpc/greet/greetpb"
	"io"
	"time"

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

	c := greetpb.NewGreetServiceClient(conn)

	// doUnary(c)

	// doStreaming(c)

	// doClientStreaming(c)

	// doBidirectional(c)

	doUnaryWithDeadline(c)
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

func doUnaryWithDeadline(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetingWithDeadlineRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Angga",
			LastName:  "Wijaya",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	res, err := c.GreetDeadline(ctx, req)
	if err != nil {
		statErr, ok := status.FromError(err)
		if ok {
			if statErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout")
			} else {
				fmt.Println("Unexpected Error", statErr)
			}
		}
	}

	if res != nil {
		fmt.Println(res.Result)
	}
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

func doBidirectional(c greetpb.GreetServiceClient) {

	req := []*greetpb.EveryoneGreetRequest{
		&greetpb.EveryoneGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Angga1",
				LastName:  "W1",
			},
		},
		&greetpb.EveryoneGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Angga2",
				LastName:  "W2",
			},
		},
		&greetpb.EveryoneGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Angga3",
				LastName:  "W3",
			},
		},
		&greetpb.EveryoneGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Angga4",
				LastName:  "W4",
			},
		},
	}

	stream, err := c.EveryoneGreet(context.Background())

	if err != nil {
		panic(err)
	}

	waitc := make(chan struct{})
	go func() {
		for _, request := range req {
			fmt.Println("Request", request)
			stream.Send(request)
			time.Sleep(100 * time.Millisecond)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
				// break
			}

			fmt.Println("Result", res.GetResult())
		}
	}()

	<-waitc

}
