package main

import (
	pb "avgRpc/proto"
	"context"
	"log"
	"time"
)

func doAvg(c pb.AvgServiceClient) {
	log.Println("doAvg was invoked")
	reqs := []*pb.AvgRequest{
		{Number: 2},
		{Number: 3},
		{Number: 5},
		{Number: 41},
	}

	stream, err := c.GetAvg(context.Background())
	if err != nil {
		log.Fatalf("Error while processing doAvg. Error: %v", err)
	}
	for _, req := range reqs {
		log.Printf("Sending %d\n", req.Number)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error on close and receiving stream, %v", err)
	}
	log.Printf("LongGreet: %f\n", res.Avg)
}
