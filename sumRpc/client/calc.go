package main

import (
	"context"
	"log"
	pb "sumRpc/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Printf("doSum function was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		Num1: 3,
		Num2: 10,
	})
	if err != nil {
		log.Fatal("Cannot use Sum service: %v\n", err)
	}

	log.Printf("Result of Sum: %s\n", res)
}
