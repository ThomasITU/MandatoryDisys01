package main

import (
	"context"
	"log"
	"os"
	"time"

	co "github.com/ThomasITU/MandatoryDisys01/course"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "courseClient"
)

func main() {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := co.NewCourseServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Contact the server and print out its response.
	request := defaultName
	if len(os.Args) > 1 {
		request = os.Args[1]
	}

	allCourses, err := c.GetCourses(ctx, &GetCoursesRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", allCourses.GetMessage())

	getCourse, err := c.GetCourseById(ctx, &GetCourseByIdRequest{Body: request})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", getCourse.GetMessage())

	deleteCourse, err := c.DeleteCourseById(ctx, &DeleteCourseByIdRequest{Body: request})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", deleteCourse.GetMessage())
}
