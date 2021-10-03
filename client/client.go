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

	// Contact the server and print out its response.
	request := defaultName
	if len(os.Args) > 1 {
		request = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	allCourses, err := c.GetCourses(ctx, &co.GetCoursesRequest{})
	if err != nil {
		log.Fatalf("could not get all courses: %v", err)
	}
	log.Printf("%s", allCourses.GetName())

	getCourse, err := c.GetCourseById(ctx, &co.GetCourseByIdRequest{Request: request})
	if err != nil {
		log.Fatalf("could not getCourse: %v", err)
	}
	log.Printf("%s", getCourse.GetName())

	deleteCourse, err := c.DeleteCourseById(ctx, &co.DeleteCourseByIdRequest{Request: request})
	if err != nil {
		log.Fatalf("could not delete: %v", err)
	}
	log.Printf("%s", deleteCourse.GetName())

	postCourse, err := c.PostCourse(ctx, &co.PostCourseRequest{Request: request})
	if err != nil {
		log.Fatalf("could not post course")
	}
	log.Printf("%s", postCourse.GetName())
}
