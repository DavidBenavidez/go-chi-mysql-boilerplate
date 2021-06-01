package course

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/davidbenavidez/go-chi-mysql-boilerplate/internal/log"
	"github.com/jinzhu/gorm"
)

func CreateCourse(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var course CreateCourseDTO

		errorDecode := json.NewDecoder(r.Body).Decode(&course)
		if errorDecode != nil {
			log.Errorf("Something went wrong decoding the body: %s", errorDecode.Error())
			respondJSON(w, nil, http.StatusInternalServerError, errorDecode)
			return
		}

		course, status, description, err := validateCreateCourse(db, course)

		response := CreateCourseResponseDTO{
			StatusCode:  status,
			Description: description,
			CourseData:  course,
		}

		respondJSON(w, response, status, err)
	}
}

func DeleteCourse(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		courseName := strings.TrimPrefix(r.URL.Path, "/course/")

		rows, status, description, err := validateDeleteCourse(db, courseName)

		response := DeleteCourseResponseDTO{
			StatusCode:   status,
			Description:  description,
			RowsAffected: rows,
		}

		respondJSON(w, response, status, err)
	}
}

func respondJSON(w http.ResponseWriter, response interface{}, status int, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
