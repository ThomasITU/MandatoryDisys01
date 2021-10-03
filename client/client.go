package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	co "github.com/ThomasITU/MandatoryDisys01/course"
	"google.golang.org/grpc"
)

var options = []string{"allCourses takes no parameters", "getCourse takes an id: \"1\"", "put takes the id of a course,workload and rating: \"1.30,50\"",
	"postCourse takes workload and rating: \"30,50\"", "delCourse takes an id: \"1\""}

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

	log.Println(optionsToString() + "\n")
	for {
		fmt.Scanln(&request)
		ctx := context.Background()

		evalRequest(request, ctx, c)
		request = ""
	}
}

//CLI logic
func evalRequest(request string, ctx context.Context, c co.CourseServiceClient) {
	if request == "options" {
		log.Println(optionsToString() + "\n")
		return
	} else if request == "allCourses" {
		allCourses(ctx, c)
		return
	}

	var parameters string
	fmt.Scanln(&parameters)

	switch request {
	case "putCourse":
		putCourse(parameters, ctx, c)
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
	log.Printf("%s\n", allCourses.GetName())
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

func putCourse(params string, ctx context.Context, c co.CourseServiceClient) {
	putCourse, err := c.PutCourse(ctx, &co.PutCourseRequest{Request: params})
	if err != nil {
		log.Fatalf("could not update course: %s", params)
	}
	log.Printf("%s", putCourse.GetName())
}

// helper method
func optionsToString() string {
	var sb strings.Builder
	for _, option := range options {
		sb.WriteString("\n" + option)
	}
	return sb.String()
}
