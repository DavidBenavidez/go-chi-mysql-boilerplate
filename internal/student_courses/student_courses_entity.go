package student_courses

// Course table
type StudentCourses struct {
	CourseName   string `gorm:"type:not null; column:course_NAME"`
	StudentEmail string `gorm:"type:not null; column:student_EMAIL"`
}
