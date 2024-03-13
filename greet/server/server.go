package main

import pb "github.com/bgunay/grpc-go-course/greet/proto"

type Server struct {
	pb.GreetServiceServer
}
