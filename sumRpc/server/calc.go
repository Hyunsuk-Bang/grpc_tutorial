package main

import (
	"context"
	"log"
	pb "sumRpc/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked")
	return &pb.SumResponse{
		Result: in.Num1 + in.Num2,
	}, nil
}
