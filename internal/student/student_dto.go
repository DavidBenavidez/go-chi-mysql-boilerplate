package student

type StudentDTO struct {
	Name    string   `json:"name" validate:"required"`
	Email   string   `json:"email" validate:"required"`
	Phone   string   `json:"phone"`
	Courses []string `json:"courses"`
}

type StudentResponseDTO struct {
	StatusCode  int        `json:"statusCode"`
	Description string     `json:"description"`
	StudentData StudentDTO `json:"student"`
}

type StudentsResponseDTO struct {
	StatusCode  int          `json:"statusCode"`
	Description string       `json:"description"`
	StudentData []StudentDTO `json:"students"`
}
