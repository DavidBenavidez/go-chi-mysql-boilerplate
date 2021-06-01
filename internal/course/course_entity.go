package course

import "github.com/davidbenavidez/go-chi-mysql-boilerplate/internal/student"

// Course table
type Course struct {
	Name          string             `gorm:"primary_key;not null;column:NAME"`
	ProfessorName string             `gorm:"not null;column:PROFESSOR_NAME"`
	Description   string             `gorm:"column:DESCRIPTION"`
	Students      []*student.Student `gorm:"many2many:student_courses;"`
}

// type Student struct {
// 	Name    string           `gorm:"not null;column:NAME"`
// 	Email   string           `gorm:"primary_key;not null;column:EMAIL"`
// 	Phone   string           `gorm:"column:PHONE"`
// 	Courses []*course.Course `gorm:"many2many:student_courses;"`
// }
