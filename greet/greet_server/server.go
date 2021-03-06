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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
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

func (s *server) GreetDeadline(ctx context.Context, req *greetpb.GreetingWithDeadlineRequest) (*greetpb.GreetingWithDeadlineResponse, error) {
	if ctx.Err() == context.Canceled {
		fmt.Println("The client canceled the request")
		return nil, status.Error(codes.Canceled, "Request Terminated")
	}
	time.Sleep(6 * time.Second)
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	res := &greetpb.GreetingWithDeadlineResponse{
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

	certFile := "certs/directoryxx.com/cert.pem"
	keyFile := "certs/directoryxx.com/privkey.pem"
	creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
	if sslErr != nil {
		fmt.Println("Failed to load : ", sslErr)
		return
	}
	opts := grpc.Creds(creds)
	s := grpc.NewServer(opts)
	greetpb.RegisterGreetServiceServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
