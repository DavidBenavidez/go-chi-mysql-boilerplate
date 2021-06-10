package student_courses

type CreateStudentCoursesDTO struct {
	StudentEmail string `json:"studentEmail" validate:"required"`
}

type CreateStudentCoursesResponseDTO struct {
	StatusCode         int                     `json:"statusCode"`
	Description        string                  `json:"description"`
	StudentCoursesData CreateStudentCoursesDTO `json:"course"`
}
