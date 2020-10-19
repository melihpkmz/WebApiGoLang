package Models



type RequestBook struct {
	Title  string    `json:"title" validate:"required"`
	Author string    `json:"author" validate:"required" `
}