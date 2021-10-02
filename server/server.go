package main

import (
	"log"
	"net"

	co "github.com/ThomasITU/MandatoryDisys01/course"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type Server struct {
	co.UnimplementedCourseServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	co.RegisterCourseServiceServer(grpcServer, &Server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %s  %v", port, err)
	}

}
