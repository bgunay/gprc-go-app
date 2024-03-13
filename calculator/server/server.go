package main

import pb "github.com/bgunay/grpc-go-course/calculator/proto"

type Server struct {
	pb.CalculatorServiceServer
}
