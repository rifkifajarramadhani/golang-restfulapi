package helper

import (
	"golang-restfulapi/entity"
)

func ToCategoryResponse(category entity.Category) entity.CategoryResponse {
	return entity.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
