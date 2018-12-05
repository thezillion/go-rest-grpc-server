package controllers

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "rest-grpc-server/app/helloworld"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)


// SayHello implements helloworld.GreeterServer
func (c *App) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	// Manipulate App context, for example
	// use c.DB... to use Database objects,
	// access Conf objects and so on.
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &App{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
