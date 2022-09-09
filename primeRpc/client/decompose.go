package main

import (
	"context"
	"io"
	"log"
	pb "primeRpc/proto"
)

func doDecompose(c pb.PrimeClient) {
	log.Println("doDecompose function was invoked")

	req := &pb.Request{
		Input: 120,
	}
	stream, err := c.Decompose(context.Background(), req)
	if err != nil {
		log.Fatalf("Err while calling decompose: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the stream %v \n")
		}
		log.Printf("%d", msg.Output)

	}
}
