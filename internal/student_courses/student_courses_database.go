package student_courses

import "github.com/jinzhu/gorm"

func insertStudentCourseDB(db *gorm.DB, studentCourse CreateStudentCoursesDTO, courseName string) (int, error) {
	// Validate parameters here
	newStudentCourse := StudentCourses{
		CourseName:   courseName,
		StudentEmail: studentCourse.StudentEmail,
	}

	result := db.Create(&newStudentCourse)

	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}
