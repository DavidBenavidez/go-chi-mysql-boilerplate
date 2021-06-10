package student_courses

import (
	"encoding/json"
	"net/http"

	"github.com/davidbenavidez/go-chi-mysql-boilerplate/internal/log"
	"github.com/davidbenavidez/go-chi-mysql-boilerplate/internal/utils"
	"github.com/jinzhu/gorm"

	"github.com/go-chi/chi"
)

func CreateStudentCourses(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var studentCourse CreateStudentCoursesDTO
		courseName := chi.URLParam(r, "courseName")

		errorDecode := json.NewDecoder(r.Body).Decode(&studentCourse)
		if errorDecode != nil {
			log.Errorf("Something went wrong decoding the body: %s", errorDecode.Error())
			utils.RespondJSON(w, nil, http.StatusInternalServerError, errorDecode)
			return
		}

		csData, status, description, err := validateCreateStudentCourse(db, studentCourse, courseName)

		response := CreateStudentCoursesResponseDTO{
			StatusCode:         status,
			Description:        description,
			StudentCoursesData: csData,
		}

		utils.RespondJSON(w, response, status, err)
	}
}
