package main

import (
	pb "avgRpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

const addr string = "0.0.0.0:50051"

type Server struct {
	pb.AvgServiceServer
}

func main() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("SERVER: cannot listen to %s", addr)
	}
	log.Printf("SERVER: Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterAvgServiceServer(s, &Server{})

	if err = s.Serve(listener); err != nil {
		log.Fatalf("Cannot serve this listener")
	}
}
