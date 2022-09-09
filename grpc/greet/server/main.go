package main

import (
	pb "grpc/greet/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{}) //grpc server need an instance for the Greet Service

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %s\n", err)
	}

}
