package course

import (
	"errors"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

func validateCreateCourse(db *gorm.DB, course CreateCourseDTO) (CreateCourseDTO, int, string, error) {
	// Validate parameters here
	if course.Name == "" {
		return CreateCourseDTO{}, http.StatusBadRequest, "Missing course name", errors.New("Missing course name")
	}

	if course.ProfessorName == "" {
		return CreateCourseDTO{}, http.StatusBadRequest, "Missing professor name", errors.New("Missing professor name")
	}

	// if all is valid, insert to database
	rowsAdded, err := insertCourseDB(db, course)

	if err != nil {
		return CreateCourseDTO{}, http.StatusInternalServerError, err.Error(), err
	}

	log.Printf("Successfully created course %+v. Rows affected: %d", course, rowsAdded)

	return course, http.StatusOK, "Successfully created course", nil
}
