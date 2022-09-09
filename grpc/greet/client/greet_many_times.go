package main

import (
	"context"
	pb "grpc/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "Hyunsuk",
	}
	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Err while calling GreetManyTimes: %v \n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the stream %v \n")
		}
		log.Printf("%s", msg.Result)
	}
}
