package course

import "github.com/davidbenavidez/go-chi-mysql-boilerplate/internal/student"

// Course table
type Course struct {
	Name          string             `gorm:"primary_key;not null;column:NAME"`
	ProfessorName string             `gorm:"not null;column:PROFESSOR_NAME"`
	Description   string             `gorm:"column:DESCRIPTION"`
	Students      []*student.Student `gorm:"many2many:student_courses;"`
}
