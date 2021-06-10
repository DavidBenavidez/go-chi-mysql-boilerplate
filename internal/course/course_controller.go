package course

import (
	"encoding/json"
	"net/http"

	"github.com/davidbenavidez/go-chi-mysql-boilerplate/internal/log"
	"github.com/davidbenavidez/go-chi-mysql-boilerplate/internal/utils"
	"github.com/jinzhu/gorm"
)

func CreateCourse(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var course CreateCourseDTO

		errorDecode := json.NewDecoder(r.Body).Decode(&course)
		if errorDecode != nil {
			log.Errorf("Something went wrong decoding the body: %s", errorDecode.Error())
			utils.RespondJSON(w, nil, http.StatusInternalServerError, errorDecode)
			return
		}

		course, status, description, err := validateCreateCourse(db, course)

		response := CreateCourseResponseDTO{
			StatusCode:  status,
			Description: description,
			CourseData:  course,
		}

		utils.RespondJSON(w, response, status, err)
	}
}
