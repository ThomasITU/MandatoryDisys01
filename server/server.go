package main

import (
	"context"
	"log"
	"net"
	"sort"
	"strconv"
	"strings"

	co "github.com/ThomasITU/MandatoryDisys01/course"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type Course struct {
	ID       string
	Workload int64
	Rating   int64
}

var courses = []Course{
	{ID: "0", Workload: 10, Rating: 80},
	{ID: "1", Workload: 10, Rating: 90},
	{ID: "2", Workload: 20, Rating: 75},
}

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

func (s *Server) GetCourses(ctx context.Context, message *co.GetCoursesRequest) (*co.Message, error) {
	log.Printf("Receive message from client: %s", message.GetRequest())
	return &co.Message{Name: "heres is all the courses:\n\n" + coursesToString()}, nil
}

func (s *Server) GetCourseById(ctx context.Context, message *co.GetCourseByIdRequest) (*co.Message, error) {
	return &co.Message{Name: "\n\nheres is the course with id: " + courseToString(message.GetRequest())}, nil
}

func (s *Server) DeleteCourseById(ctx context.Context, message *co.DeleteCourseByIdRequest) (*co.Message, error) {
	deletionComplete := deleteCourseByID(message.GetRequest())
	return &co.Message{Name: deletionComplete}, nil
}

// helper method

func deleteCourseByID(id string) string {
	deletionState := "Deletion failed couldn't find and delete: " + id
	for _, course := range courses {
		if course.ID == id {
			deletionState = courseToString(id) + "\n has been deleted"
			delCourse(id)
		}
	}
	return deletionState
}

func courseToString(id string) string {
	courseString := "course not found"
	for _, course := range courses {
		if course.ID == id {
			courseString = course.ID + " the course workload is: " + strconv.FormatInt(course.Workload, 10) + " and is rated: " + strconv.FormatInt(course.Rating, 10)
		}
	}
	return courseString
}

func coursesToString() string {
	var sb strings.Builder
	for _, course := range courses {
		sb.WriteString("courseid: " + course.ID + " has a workload of: ")
		sb.WriteString(strconv.FormatInt(course.Workload, 10))
		sb.WriteString(" and is rated: " + strconv.FormatInt(course.Rating, 10) + "\n")
	}
	return sb.String()
}

// copy pasta fra Lotte's git
func delCourse(id string) {
	oldCourses := courses
	courses = courses[0:0]

	for _, a := range oldCourses {
		ID := a.ID
		if ID == id {
			for _, c := range oldCourses {
				if c.ID != a.ID {
					courses = append(courses, c)
				}
			}
		}
	}

	sort.Slice(courses, func(i, j int) bool {
		return courses[i].ID < courses[j].ID
	})

}
