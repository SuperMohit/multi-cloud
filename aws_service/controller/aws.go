package controller

import (
	"context"
	"log"
	pb "github.com/aws_service/protogen"
)



// server is used to implement helloworld.GreeterServer.
type AWSController struct {
	//inject services
	pb.UnimplementedAWSServer
}

type CloudController interface {
	SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error)
}


// SayHello implements helloworld.GreeterServer
func (s *AWSController) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

