package course_student

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/davidbenavidez/go-chi-mysql-boilerplate/internal/log"
	"github.com/jinzhu/gorm"
)

func CreateCourseStudent(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var courseStudent CreateCourseStudentDTO
		courseNameNoPrefix := strings.TrimPrefix(r.URL.Path, "/courses/")
		courseName := strings.TrimSuffix(courseNameNoPrefix, "/students")

		errorDecode := json.NewDecoder(r.Body).Decode(&courseStudent)
		if errorDecode != nil {
			log.Errorf("Something went wrong decoding the body: %s", errorDecode.Error())
			respondJSON(w, nil, http.StatusInternalServerError, errorDecode)
			return
		}

		csData, status, description, err := validateCreateCourseStudent(db, courseStudent, courseName)

		response := CreateCourseStudentResponseDTO{
			StatusCode:        status,
			Description:       description,
			CourseStudentData: csData,
		}

		respondJSON(w, response, status, err)
	}
}

func GetCourseStudent(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 		someRegex, regexErr := regexp.Compile("/course/(.*)/student/(.*)")
		// 		if regexErr != nil {
		// 			respondJSON(w, nil, http.StatusInternalServerError, regexErr)
		// 		}

		// 		matches := someRegex.FindStringSubmatch(r.URL.Path)
		// 		courseName := matches[1]
		// 		studentName := matches[2]

		// 		studentData, status, description, err := validateGetCourseStudent(db, courseName, studentName)

		// 		response := CreateCourseStudentResponseDTO{
		// 			StatusCode:  status,
		// 			Description: description,
		// 			StudentData: studentData,
		// 		}

		// 		respondJSON(w, response, status, err)
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
