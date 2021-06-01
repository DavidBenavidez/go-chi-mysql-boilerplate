package student

import (
	"encoding/json"
	"net/http"

	"github.com/davidbenavidez/go-chi-mysql-boilerplate/internal/log"
	"github.com/jinzhu/gorm"
)

func CreateStudent(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student CreateStudentDTO

		errorDecode := json.NewDecoder(r.Body).Decode(&student)
		if errorDecode != nil {
			log.Errorf("Something went wrong decoding the body: %s", errorDecode.Error())
			respondJSON(w, nil, http.StatusInternalServerError, errorDecode)
			return
		}

		student, status, description, err := validateCreateStudent(db, student)

		response := CreateStudentResponseDTO{
			StatusCode:  status,
			Description: description,
			StudentData: student,
		}

		respondJSON(w, response, status, err)
	}
}

func GetStudents(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		students, status, description, err := validateGetStudents(db)

		response := GetStudentsResponseDTO{
			StatusCode:  status,
			Description: description,
			StudentData: students,
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
