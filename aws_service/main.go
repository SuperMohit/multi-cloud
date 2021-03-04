package main

import (
	"google.golang.org/grpc"
	ctrl "github.com/aws_service/controller"
	"log"
	"net"
	pb "github.com/aws_service/protogen"
)

const (
	port = ":50051"
)


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAWSServer(s, &ctrl.AWSController{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}