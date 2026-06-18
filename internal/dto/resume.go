package dto

type CreateResumeRequest struct {
	Title  string `json:"title" validate:"required,min=3,max=100"`
	About  string `json:"about" validate:"required,min=10,max=5000"`
	Skills string `json:"skills" validate:"required,min=3,max=1000"`
}

type UpdateResumeRequest struct {
	Id     uint   `json:"id"`
	Title  string `json:"title" validate:"required,min=3,max=100"`
	About  string `json:"about" validate:"required,min=10,max=5000"`
	Skills string `json:"skills" validate:"required,min=3,max=1000"`
}

type ResumeResponse struct {
	ID uint `json:"id"`

	Title  string       `json:"title"`
	About  string       `json:"about"`
	Skills []string     `json:"skills"`
	User   UserResponse `json:"user"`
}
