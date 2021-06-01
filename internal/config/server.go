package config

import (
	"net/http"

	"github.com/davidbenavidez/go-chi-mysql-boilerplate/internal/course"
	cs "github.com/davidbenavidez/go-chi-mysql-boilerplate/internal/course_student"
	"github.com/davidbenavidez/go-chi-mysql-boilerplate/internal/student"
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

type Server struct {
	router *chi.Mux
	db     *gorm.DB
	Port   string
}

type Config struct {
	Server   Server
	Database DBConfiguration
}

var serverConfig *Config

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// Course Handlers
func (s *Server) handleCreateCourse() http.HandlerFunc {
	return course.CreateCourse(s.db)
}
func (s *Server) handleDeleteCourse() http.HandlerFunc {
	return course.DeleteCourse(s.db)
}

// Student Handlers
func (s *Server) handleCreateStudent() http.HandlerFunc {
	return student.CreateStudent(s.db)
}

func (s *Server) handleGetStudents() http.HandlerFunc {
	return student.GetStudents(s.db)
}

// Course Student Handlers
func (s *Server) handleCreateCourseStudent() http.HandlerFunc {
	return cs.CreateCourseStudent(s.db)
}
func (s *Server) handleGetCourseStudent() http.HandlerFunc {
	return cs.GetCourseStudent(s.db)
}

func SetupServer() (*Server, string, error) {
	var err error
	s := &Server{}

	// Setup DB connection
	s.db, err = s.setupDatabase()

	if err != nil {
		return nil, "", err
	}

	// inject handlers to routes
	s.setupRoutes()

	port := ":" + serverConfig.Server.Port
	return s, port, nil
}
