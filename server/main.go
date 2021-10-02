package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"../mandatory exercise 1/mandatoryExercise/course"
)

const (
	port = ":50051"
)

type server struct {
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := course.Server{}

	grpcServer := grpc.NewServer()

	course.RegisterChatServiceServer(grpcServer, &s)
	if err := grpcServer.Server(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %s  %v", port, err)
	}

}
