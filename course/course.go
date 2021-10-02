package course

import (
	"log"
	"sort"
	"strconv"
	"strings"

	"context"
)

type Server struct {
	UnimplementedCourseServiceServer
}

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

func (s *Server) GetCourses(ctx context.Context, message *GetCoursesRequest) (*Message, error) {
	log.Printf("Receive message from client: %s", message.GetRequest())
	return &Message{Name: "heres is all the courses:\n\n" + coursesToString()}, nil
}

func (s *Server) GetCourseById(ctx context.Context, message *GetCourseByIdRequest) (*Message, error) {
	return &Message{Name: "\n\nheres is the course with id: " + courseToString(message.GetRequest())}, nil
}

func (s *Server) DeleteCourseById(ctx context.Context, message *DeleteCourseByIdRequest) (*Message, error) {
	deletionComplete := deleteCourseByID(message.GetRequest())
	return &Message{Name: deletionComplete}, nil
}

// func (s *server) PutCourseById(ctx context.Context, in *pb.PutCourseByIdRequest) (*pb.PutCourseByIdReply, error) {

// }

// func (s *server) PostCourse(ctx context.Context, in *pb.PostCourseRequest) (*pb.PostCourseReply, error) {

// }

// void post(Course c)
// void put(course c)

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
