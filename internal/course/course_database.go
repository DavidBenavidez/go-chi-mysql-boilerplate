package course

import (
	"github.com/jinzhu/gorm"
)

func insertCourseDB(db *gorm.DB, course CreateCourseDTO) (int, error) {
	// Validate parameters here
	newCourse := Course{
		Name:          course.Name,
		ProfessorName: course.ProfessorName,
		Description:   course.Description,
	}

	result := db.Create(&newCourse)

	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}
