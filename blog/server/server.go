package main

import pb "github.com/bgunay/grpc-go-course/blog/proto"

type Server struct {
	pb.BlogServiceServer
}
