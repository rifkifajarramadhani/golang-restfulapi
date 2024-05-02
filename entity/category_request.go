package entity

type CategoryCreateRequest struct {
	Name string
}

type CategoryUpdateRequest struct {
	Id   int
	Name string
}
