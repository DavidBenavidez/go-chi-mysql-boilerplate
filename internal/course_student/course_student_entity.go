package course_student

import (
	c "github.com/davidbenavidez/go-chi-mysql-boilerplate/internal/course"
	s "github.com/davidbenavidez/go-chi-mysql-boilerplate/internal/student"
)

// Course table
type CourseStudent struct {
	CourseRefer  string    `gorm:"primary_key;not null;column:COURSE_NAME"`
	StudentRefer string    `gorm:"primary_key;not null;column:STUDENT_NAME"`
	course       c.Course  `gorm:"foreignKey:CourseRefer;constraint:OnDelete:SET NULL;"`
	student      s.Student `gorm:"foreignKey:studentRefer;constraint:OnDelete:SET NULL;"`
}
