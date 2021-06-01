package student

// Student table
// type Student struct {
// Name  string `gorm:"not null;column:NAME"`
// Email string `gorm:"primary_key;not null;column:EMAIL"`
// Phone string `gorm:"column:PHONE"`
// }

type CourseStudent struct {
	CourseName  string `gorm:"column:COURSE_NAME"`
	StudentName string `gorm:"column:STUDENT_NAME"`
}

type Student struct {
	Name    string    `gorm:"not null;column:NAME"`
	Email   string    `gorm:"primary_key;not null;column:EMAIL"`
	Phone   string    `gorm:"column:PHONE"`
	Courses []*Course `gorm:"many2many:student_courses;"`
}

type Course struct {
	Name          string     `gorm:"primary_key;not null;column:NAME"`
	ProfessorName string     `gorm:"not null;column:PROFESSOR_NAME"`
	Description   string     `gorm:"column:DESCRIPTION"`
	Students      []*Student `gorm:"many2many:student_courses;"`
}
