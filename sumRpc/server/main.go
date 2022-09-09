package main

import (
	"log"
	"net"
	pb "sumRpc/proto"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("connection fail %v\n", err)
	}
	log.Printf("SERVER: Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err = s.Serve(listener); err != nil {
		log.Fatalf("fail to serve %s\n", addr)
	}
}
