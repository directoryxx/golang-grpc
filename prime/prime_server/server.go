package main

import (
	"fmt"
	"golang-grpc/prime/protoprime"
	"net"

	"google.golang.org/grpc"
)

type serverprime struct{}

func (s *serverprime) Prime(req *protoprime.PrimeRequest, stream protoprime.PrimeService_PrimeServer) error {
	a := req.GetPrime().A
	k := 2
	for a > 1 {
		if a%int64(k) == 0 {
			res := &protoprime.PrimeResponse{
				Result: int64(k),
			}
			stream.Send(res)
			a = a / int64(k)
		} else {
			k++
		}

	}

	return nil

}

func main() {
	fmt.Println("Prime server started")
	lis, err := net.Listen("tcp", "0.0.0.0:8010")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	protoprime.RegisterPrimeServiceServer(s, &serverprime{})

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
