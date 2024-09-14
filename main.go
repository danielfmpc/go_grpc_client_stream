package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/danielfmpc/client_go_rpc_client_stream/src/pb/calc"
	"google.golang.org/grpc"
)

type Server struct {
	calc.CalcServiceServer
}

func (s *Server) Calc(stream calc.CalcService_CalcServer) error {
	var quantity int32 = 0
	var total int32 = 0
	for {
		input, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&calc.Output{
				Quantity: quantity,
				Avarage:  float64(float64(total) / float64(quantity)),
				Total:    total,
			})
		}
		if err != nil {
			return err
		}

		quantity++
		total += input.GetValue()

		fmt.Printf("input %+v\n", input)
	}
}

func main() {
	log.Println("Start server")
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln("err listerner", err)
	}

	s := grpc.NewServer()
	calc.RegisterCalcServiceServer(s, &Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalln("err server", err)
	}
}
