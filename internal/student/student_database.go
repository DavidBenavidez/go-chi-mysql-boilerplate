package student

import (
	"github.com/jinzhu/gorm"
)

func insertStudentDB(db *gorm.DB, student StudentDTO) (int, error) {
	// Validate parameters here
	newStudent := Student{
		Name:  student.Name,
		Email: student.Email,
		Phone: student.Phone,
	}

	result := db.Create(&newStudent)

	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

func getStudentsDB(db *gorm.DB) ([]StudentDTO, error) {
	// Validate parameters here
	var studentsList []StudentDTO
	var students []Student
	result := db.Table("student").Find(&students)

	if result.Error != nil {
		return []StudentDTO{}, result.Error
	}

	for _, student := range students {
		var courses []string
		var studentCourses []StudentCourses

		db.Table("student_courses").Where("student_EMAIL = ?", student.Email).Find(&studentCourses)

		for _, course := range studentCourses {
			courses = append(courses, course.CourseName)
		}

		studentsList = append(studentsList, StudentDTO{
			Name:    student.Name,
			Email:   student.Email,
			Phone:   student.Phone,
			Courses: courses,
		})
	}

	return studentsList, nil
}
