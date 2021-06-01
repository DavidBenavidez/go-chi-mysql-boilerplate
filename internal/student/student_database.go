package student

import (
	"github.com/jinzhu/gorm"
)

func insertStudentDB(db *gorm.DB, student CreateStudentDTO) (int, error) {
	// Validate parameters here
	newCourse := Student(student)

	result := db.Create(&newCourse)

	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

// func getStudentsDB(db *gorm.DB) ([]StudentDTO, error) {
func getStudentsDB(db *gorm.DB) ([]StudentDTO, error) {
	// Validate parameters here
	var students []Student
	var studentsDTO []StudentDTO
	result := db.Table("student").Find(&students)

	if result.Error != nil {
		return []StudentDTO{}, result.Error
	}

	for _, student := range students {
		var csData []CourseStudent

		db.Model(&CourseStudent{}).Select("course_student.")
		resultCS := db.Debug().Table("course_student").Where("STUDENT_NAME = ?", student.Name).Find(&csData)
		if resultCS.Error != nil {
			return []StudentDTO{}, resultCS.Error
		}

		var courses []string
		for _, csd := range csData {
			courses = append(courses, csd.CourseName)
		}

		newStudentDTO := StudentDTO{
			Name:    student.Name,
			Email:   student.Email,
			Phone:   student.Phone,
			Courses: courses,
		}

		studentsDTO = append(studentsDTO, newStudentDTO)
	}

	return studentsDTO, nil
}
