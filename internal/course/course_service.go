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

func validateDeleteCourse(db *gorm.DB, courseName string) (int, int, string, error) {
	// Validate parameters here
	if courseName == "" {
		return 0, http.StatusBadRequest, "Missing course name", errors.New("Missing course name")
	}

	// if all is valid, delete from database
	rowsAffected, err := deleteCourseDB(db, courseName)

	if err != nil {
		return 0, http.StatusInternalServerError, err.Error(), err
	}

	if rowsAffected == 0 {
		return 0, http.StatusBadRequest, "Course not found", errors.New("Course not found")
	}

	log.Printf("Successfully deleted course. Rows affected: %d", rowsAffected)

	return rowsAffected, http.StatusOK, "Successfully deleted course", nil
}
