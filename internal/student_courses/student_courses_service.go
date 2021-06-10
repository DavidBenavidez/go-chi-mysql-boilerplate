package student_courses

import (
	"errors"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

func validateCreateStudentCourse(db *gorm.DB, studentCourse CreateStudentCoursesDTO, courseName string) (CreateStudentCoursesDTO, int, string, error) {
	// Validate parameters here
	if studentCourse.StudentEmail == "" {
		return CreateStudentCoursesDTO{}, http.StatusBadRequest, "Missing course name", errors.New("Missing course name")
	}

	if courseName == "" {
		return CreateStudentCoursesDTO{}, http.StatusBadRequest, "Missing student name", errors.New("Missing student name")
	}

	// if all is valid, insert to database
	rowsAdded, err := insertStudentCourseDB(db, studentCourse, courseName)

	if err != nil {
		return CreateStudentCoursesDTO{}, http.StatusInternalServerError, err.Error(), err
	}

	log.Printf("Successfully created course %+v. Rows affected: %d", studentCourse, rowsAdded)

	return studentCourse, http.StatusOK, "Successfully added student to course", nil
}
