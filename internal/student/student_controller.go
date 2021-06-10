package student

import (
	"encoding/json"
	"net/http"

	"github.com/davidbenavidez/go-chi-mysql-boilerplate/internal/log"
	"github.com/davidbenavidez/go-chi-mysql-boilerplate/internal/utils"
	"github.com/jinzhu/gorm"
)

func CreateStudent(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student StudentDTO

		errorDecode := json.NewDecoder(r.Body).Decode(&student)
		if errorDecode != nil {
			log.Errorf("Something went wrong decoding the body: %s", errorDecode.Error())

			utils.RespondJSON(w, nil, http.StatusInternalServerError, errorDecode)
			return
		}

		student, status, description, err := validateCreateStudent(db, student)

		response := StudentResponseDTO{
			StatusCode:  status,
			Description: description,
			StudentData: student,
		}

		utils.RespondJSON(w, response, status, err)
	}
}

func GetStudents(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		students, status, description, err := validateGetStudents(db)

		response := StudentsResponseDTO{
			StatusCode:  status,
			Description: description,
			StudentData: students,
		}

		utils.RespondJSON(w, response, status, err)
	}
}
