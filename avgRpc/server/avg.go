package main

import (
	pb "avgRpc/proto"
	"io"
	"log"
)

func (s *Server) GetAvg(stream pb.AvgService_GetAvgServer) error {
	var ret float64 = 0.0
	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Avg: ret / float64(count),
			})
		}
		if err != nil {
			log.Fatal("Error while reading client stream %v\n", err)
		}
		ret += float64(req.Number)
		count++
	}
}
