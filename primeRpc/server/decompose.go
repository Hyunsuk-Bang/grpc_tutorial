package main

import (
	pb "primeRpc/proto"
)

func (s *Server) Decompose(in *pb.Request, stream pb.Prime_DecomposeServer) error {
	var k uint64 = 2
	n := in.Input
	for n > 1 {
		if n%k == 0 {
			stream.Send(&pb.Response{
				Output: k,
			})
			n = n / k
		} else {
			k++
		}
	}
	return nil
}
