package main

import (
	"context"
	"fmt"
	"log"

	co "github.com/ThomasITU/MandatoryDisys01/course"
	"google.golang.org/grpc"
)

const (
	address        = "localhost:50051"
	defaultRequest = "allCourses"
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
	request := defaultRequest

	for {
		fmt.Scanln(&request)
		ctx := context.Background()

		evalRequest(request, ctx, c)
		request = ""
	}
}

func evalRequest(request string, ctx context.Context, c co.CourseServiceClient) {
	var parameters string
	fmt.Scanln(&parameters)

	switch request {
	case "put":
		put(parameters, ctx, c)
	case "allCourses":
		allCourses(ctx, c)
	case "getCourse":
		getCourse(parameters, ctx, c)
	case "deleteCourse":
		deleteCourse(parameters, ctx, c)
	case "postCourse":
		postCourse(parameters, ctx, c)
	default:
		break
	}
}

func allCourses(ctx context.Context, c co.CourseServiceClient) {
	allCourses, err := c.GetCourses(ctx, &co.GetCoursesRequest{})
	if err != nil {
		log.Fatalf("could not get all courses: %v", err)
	}
	log.Printf("%s", allCourses.GetName())
}

func getCourse(params string, ctx context.Context, c co.CourseServiceClient) {
	getCourse, err := c.GetCourseById(ctx, &co.GetCourseByIdRequest{Request: params})
	if err != nil {
		log.Fatalf("could not getCourse: %s", params)
	}
	log.Printf("%s", getCourse.GetName())
}

func deleteCourse(params string, ctx context.Context, c co.CourseServiceClient) {
	deleteCourse, err := c.DeleteCourseById(ctx, &co.DeleteCourseByIdRequest{Request: params})
	if err != nil {
		log.Fatalf("could not delete: %s", params)
	}
	log.Printf("%s", deleteCourse.GetName())
}

func postCourse(params string, ctx context.Context, c co.CourseServiceClient) {
	postCourse, err := c.PostCourse(ctx, &co.PostCourseRequest{Request: params})
	if err != nil {
		log.Fatalf("could not post course %s", params)
	}
	log.Printf("%s", postCourse.GetName())
}

func put(params string, ctx context.Context, c co.CourseServiceClient) {
	putCourse, err := c.PutCourse(ctx, &co.PutCourseRequest{Request: params})
	if err != nil {
		log.Fatalf("could not update course: %s", params)
	}
	log.Printf("%s", putCourse.GetName())
}
