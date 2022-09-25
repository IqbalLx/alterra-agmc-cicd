package entities

type Book struct {
	Id     uint   `json:"id"`
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
}

type UpdateBookDTO struct {
	Id     uint   `param:"id" validate:"required"`
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
}
