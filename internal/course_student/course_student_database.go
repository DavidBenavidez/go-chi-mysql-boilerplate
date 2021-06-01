package course_student

import "github.com/jinzhu/gorm"

// import (
// 	"github.com/jinzhu/gorm"
// )

func insertCourseStudentDB(db *gorm.DB, courseStudent CreateCourseStudentDTO, courseName string) (int, error) {
	// Validate parameters here
	newCourseStudent := CourseStudent{
		CourseRefer:  courseName,
		StudentRefer: courseStudent.StudentName,
	}

	result := db.Create(&newCourseStudent)

	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

// func deleteCourseDB(db *gorm.DB, name string) (int, error) {
// 	// Validate parameters here
// 	result := db.Debug().Delete(&Course{}, "NAME = ?", name)

// 	if result.Error != nil {
// 		return 0, result.Error
// 	}

// 	return int(result.RowsAffected), nil
// }
