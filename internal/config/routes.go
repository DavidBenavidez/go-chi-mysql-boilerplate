package config

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (s *Server) setupRoutes() {
	s.router = chi.NewRouter()
	s.router.Use(middleware.Logger)
	// Course
	s.router.Post("/courses", s.handleCreateCourse())
	s.router.Delete("/courses/{courseName}", s.handleDeleteCourse())

	// Student
	s.router.Get("/students", s.handleGetStudents())
	s.router.Post("/students", s.handleCreateStudent())

	// course_students
	s.router.Post("/courses/{courseName}/students", s.handleCreateCourseStudent())
	s.router.Get("/courses/{courseName}/students/{studentName}", s.handleGetCourseStudent())
}
