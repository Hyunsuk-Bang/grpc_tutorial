package main

import (
	"log"
	pb "primeRpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("CLIENT: Cannot Dial to grpc server, %v", err)
	}
	defer conn.Close()

	c := pb.NewPrimeClient(conn)
	doDecompose(c)
}
