package student

import (
	"errors"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

func validateCreateStudent(db *gorm.DB, student CreateStudentDTO) (CreateStudentDTO, int, string, error) {
	// Validate parameters here
	if student.Name == "" {
		return CreateStudentDTO{}, http.StatusBadRequest, "Missing student name", errors.New("Missing student name")
	}

	if student.Email == "" {
		return CreateStudentDTO{}, http.StatusBadRequest, "Missing student email", errors.New("Missing email")
	}

	// if all is valid, insert to database
	rowsAdded, err := insertStudentDB(db, student)

	if err != nil {
		return CreateStudentDTO{}, http.StatusInternalServerError, err.Error(), err
	}

	log.Printf("Successfully created student %+v. Rows affected: %d", student, rowsAdded)

	return student, http.StatusOK, "Successfully created student", nil
}

func validateGetStudents(db *gorm.DB) ([]StudentDTO, int, string, error) {
	// if all is valid, insert to database
	students, err := getStudentsDB(db)

	if err != nil {
		return []StudentDTO{}, http.StatusInternalServerError, err.Error(), err
	}

	return students, http.StatusOK, "Successfully retrieved students", nil
}
