package main

import (
	"context"
	"fmt"
	"golang-grpc/greet/greetpb"
	"io"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetingRequest) (*greetpb.GreetingResponse, error) {
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	res := &greetpb.GreetingResponse{
		Result: "Hello " + firstName + " " + lastName,
	}

	return res, nil
}

func (s *server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Println("Get request : ", &req)
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " " + lastName + " : " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1 * time.Second)
	}

	return nil
}

func (s *server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	result := "Hello "
	for {
		req, err := stream.Recv()
		firstName := req.GetGreeting().GetFirstName()
		lastName := req.GetGreeting().GetLastName()
		result += firstName + " " + lastName + " !"
		if err == io.EOF {
			res := &greetpb.LongGreetResponse{
				Result: result,
			}
			// Sudah paling terakhir (End of file)
			// Kirim balik datanya
			return stream.SendAndClose(res)
		}
		if err != nil {
			panic(err)
		}

	}
}

func (s *server) EveryoneGreet(stream greetpb.GreetService_EveryoneGreetServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			panic(err)
		}
		firstName := req.GetGreeting().GetFirstName()
		lastName := req.GetGreeting().GetLastName()
		result := firstName + " " + lastName + " !"
		sendErr := stream.Send(&greetpb.EveryoneGreetResponse{
			Result: result,
		})

		if sendErr != nil {
			panic(sendErr)
		}

	}
}

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
