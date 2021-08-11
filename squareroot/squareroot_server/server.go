package main

import (
	"context"
	"fmt"
	"math"
	"net"

	"golang-grpc/squareroot/protosquareroot"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type squarerootserver struct{}

func (s *squarerootserver) SquareRoot(ctx context.Context, req *protosquareroot.SquareRootRequest) (*protosquareroot.SquareRootResponse, error) {
	a := req.GetNumber()
	if a < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received negative number %v", a),
		)
	}
	res := &protosquareroot.SquareRootResponse{
		Result: int64(math.Sqrt(float64(a))),
	}

	return res, nil
}

func main() {
	fmt.Println("Squareroot server started")
	lis, err := net.Listen("tcp", "0.0.0.0:8010")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	protosquareroot.RegisterPrimeServiceServer(s, &squarerootserver{})

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
