package main

import (
	"log"
	"net"
	pb "primeRpc/proto"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.PrimeServer
}

func main() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("SERVER: %v", err)
	}
	log.Printf("Listening on %s", addr)

	s := grpc.NewServer()
	pb.RegisterPrimeServer(s, &Server{})

	if err = s.Serve(listener); err != nil {
		log.Fatalf("SERVER: %v", err)
	}
}
