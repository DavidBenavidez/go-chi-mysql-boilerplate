package course_student

import (
	"errors"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

func validateCreateCourseStudent(db *gorm.DB, courseStudent CreateCourseStudentDTO, courseName string) (CreateCourseStudentDTO, int, string, error) {
	// Validate parameters here
	if courseStudent.StudentName == "" {
		return CreateCourseStudentDTO{}, http.StatusBadRequest, "Missing course name", errors.New("Missing course name")
	}

	if courseName == "" {
		return CreateCourseStudentDTO{}, http.StatusBadRequest, "Missing student name", errors.New("Missing student name")
	}

	// if all is valid, insert to database
	rowsAdded, err := insertCourseStudentDB(db, courseStudent, courseName)

	if err != nil {
		return CreateCourseStudentDTO{}, http.StatusInternalServerError, err.Error(), err
	}

	log.Printf("Successfully created course %+v. Rows affected: %d", courseStudent, rowsAdded)

	return courseStudent, http.StatusOK, "Successfully added student to course", nil
}

// func validateDeleteCourse(db *gorm.DB, courseName string) (int, int, string, error) {
// 	// Validate parameters here
// 	if courseName == "" {
// 		return 0, http.StatusBadRequest, "Missing course name", errors.New("Missing course name")
// 	}

// 	// if all is valid, delete from database
// 	rowsAffected, err := deleteCourseDB(db, courseName)

// 	if err != nil {
// 		return 0, http.StatusInternalServerError, err.Error(), err
// 	}

// 	if rowsAffected == 0 {
// 		return 0, http.StatusBadRequest, "Course not found", errors.New("Course not found")
// 	}

// 	log.Printf("Successfully deleted course. Rows affected: %d", rowsAffected)

// 	return rowsAffected, http.StatusOK, "Successfully deleted course", nil
// }
