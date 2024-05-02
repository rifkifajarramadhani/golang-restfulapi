package service

import (
	"context"
	"database/sql"
	"golang-restfulapi/entity"
	"golang-restfulapi/helper"
	"golang-restfulapi/repository"
)

type ICategoryService interface {
	Create(ctx context.Context, req entity.CategoryCreateRequest) entity.CategoryResponse
	Update(ctx context.Context, req entity.CategoryUpdateRequest) entity.CategoryResponse
	Delete(ctx context.Context, id int)
	Get(ctx context.Context) []entity.CategoryResponse
	GetById(ctx context.Context, id int) entity.CategoryResponse
}

type CategoryService struct {
	categoryRepository repository.CategoryRepository
	DB                 *sql.DB
}

func (s *CategoryService) Create(ctx context.Context, req entity.CategoryCreateRequest) entity.CategoryResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := entity.Category{
		Name: req.Name,
	}

	category = s.categoryRepository.Create(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (s *CategoryService) Update(ctx context.Context, req entity.CategoryUpdateRequest) entity.CategoryResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := s.categoryRepository.GetById(ctx, tx, req.Id)
	helper.PanicIfError(err)

	category.Name = req.Name

	category = s.categoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (s *CategoryService) Delete(ctx context.Context, id int) {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := s.categoryRepository.GetById(ctx, tx, id)
	helper.PanicIfError(err)

	s.categoryRepository.Delete(ctx, tx, category)
}

func (s *CategoryService) Get(ctx context.Context) []entity.CategoryResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := s.categoryRepository.Get(ctx, tx)

	categoryResponses := []entity.CategoryResponse{}
	for _, category := range categories {
		categoryResponses = append(categoryResponses, helper.ToCategoryResponse(category))
	}

	return categoryResponses
}

func (s *CategoryService) GetById(ctx context.Context, id int) entity.CategoryResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := s.categoryRepository.GetById(ctx, tx, id)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}
