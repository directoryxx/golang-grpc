package main

import (
	"context"
	"fmt"
	"net"

	"golang-grpc/sum/protosum"

	"google.golang.org/grpc"
)

type serversum struct{}

func (s *serversum) Sum(ctx context.Context, req *protosum.SumRequest) (*protosum.SumResponse, error) {
	a := req.GetSum().GetA()
	b := req.GetSum().GetB()
	result := a + b
	res := &protosum.SumResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	fmt.Println("Sum Server running.....")
	lis, err := net.Listen("tcp", "0.0.0.0:8010")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	protosum.RegisterSumServiceServer(s, &serversum{})

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
