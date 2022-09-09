package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "sumRpc/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("CLIENT: Cannot connect to the host")
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)
	doSum(c)
}
