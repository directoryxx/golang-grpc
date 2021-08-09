package main

import (
	"fmt"
	"io"
	"net"

	"golang-grpc/average/protoaverage"

	"google.golang.org/grpc"
)

type serveravg struct{}

func (s *serveravg) Average(stream protoaverage.AverageService_AverageServer) error {
	result := 0
	index := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Println(result)
			fmt.Println(index)
			avg := float32(result) / float32(index)
			res := &protoaverage.AverageResponse{
				Result: avg,
			}
			// Sudah paling terakhir (End of file)
			// Kirim balik datanya
			return stream.SendAndClose(res)
		}
		if err != nil {
			panic(err)
		}
		a := req.GetAverage().GetA()
		// lastName := req.GetGreeting().GetLastName()
		result += int(a)
		index++

	}
}

func main() {
	fmt.Println("Average Server running.....")
	lis, err := net.Listen("tcp", "0.0.0.0:8010")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	protoaverage.RegisterAverageServiceServer(s, &serveravg{})

	if err := s.Serve(lis); err != nil {
		panic(err)
	}

}
