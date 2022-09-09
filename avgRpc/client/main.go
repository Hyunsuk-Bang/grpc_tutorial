package main

import (
	"log"

	pb "avgRpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot Dial to %s. Error: %v", addr, err)
	}
	defer conn.Close()

	c := pb.NewAvgServiceClient(conn)
	doAvg(c)
}
