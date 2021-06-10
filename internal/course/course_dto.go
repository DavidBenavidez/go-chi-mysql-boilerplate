package course

type CreateCourseDTO struct {
	Name          string `json:"name" validate:"required"`
	ProfessorName string `json:"professorName" validate:"required"`
	Description   string `json:"description"`
}

type CreateCourseResponseDTO struct {
	StatusCode  int             `json:"statusCode"`
	Description string          `json:"description"`
	CourseData  CreateCourseDTO `json:"course"`
}
