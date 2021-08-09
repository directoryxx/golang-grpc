package main

import (
	"fmt"
	"io"
	"net"

	"golang-grpc/findmax/protofindmax"

	"google.golang.org/grpc"
)

type serverfindmax struct{}

func (s *serverfindmax) Findmax(stream protofindmax.FindmaxService_FindmaxServer) error {
	maximum := int64(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			panic(err)
		}
		a := req.GetFindmax().GetA()

		if maximum < a {
			maximum = a
			sendErr := stream.Send(&protofindmax.FindmaxResponse{
				Result: maximum,
			})

			if sendErr != nil {
				panic(sendErr)
			}
		}

	}
}

func main() {
	fmt.Println("FindMax Server running.....")
	lis, err := net.Listen("tcp", "0.0.0.0:8010")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	protofindmax.RegisterFindmaxServiceServer(s, &serverfindmax{})

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
