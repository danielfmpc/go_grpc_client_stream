package main

import (
	"context"
	"fmt"
	"log"

	"github.com/danielfmpc/client/client_go_rpc_client_stream/src/pb/calc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("err ", err)
	}

	defer conn.Close()

	client := calc.NewCalcServiceClient(conn)

	stream, err := client.Calc(context.Background())
	if err != nil {
		log.Fatalf("", err)
	}

	nums := []int32{2, 3, 4, 1}
	for _, v := range nums {
		if err := stream.Send(&calc.Input{Value: v}); err != nil {
			log.Fatalf("erro ", err)
		}
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("erro", err)
	}

	fmt.Println("response ", response)
}
