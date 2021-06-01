package course_student

type CreateCourseStudentDTO struct {
	StudentName string `json:"studentName" validate:"required"`
}

type CreateCourseStudentResponseDTO struct {
	StatusCode        int                    `json:"statusCode"`
	Description       string                 `json:"description"`
	CourseStudentData CreateCourseStudentDTO `json:"course"`
}
