package main

import (
	"context"
	pb "grpc/greet/proto"
	"log"
	"time"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")
	reqs := []*pb.GreetRequest{
		{FirstName: "Hyunsuk"},
		{FirstName: "James"},
		{FirstName: "Junewon"},
		{FirstName: "Brad"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while callin LongGreet: %v", err)
	}
	for _, req := range reqs {
		log.Printf("Sending %s\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving respoinse from LongGreet %v\n", err)
	}
	log.Printf("LongGreet: %s\n", res.Result)
}
