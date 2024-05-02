package entity

type CategoryCreateRequest struct {
	Name string `validate:"required,min=1,max=200"`
}

type CategoryUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,min=1,,max=200"`
}
