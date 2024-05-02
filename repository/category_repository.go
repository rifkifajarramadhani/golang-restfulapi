package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-restfulapi/entity"
	"golang-restfulapi/helper"
)

type ICategoryRepository interface {
	Create(ctx context.Context, tx *sql.Tx, c entity.Category) entity.Category
	Update(ctx context.Context, tx *sql.Tx, c entity.Category) entity.Category
	Delete(ctx context.Context, tx *sql.Tx, c entity.Category)
	Get(ctx context.Context, tx *sql.Tx) []entity.Category
	GetById(ctx context.Context, tx *sql.Tx, id int) (entity.Category, error)
}

type CategoryRepository struct {
}

func (r *CategoryRepository) Create(ctx context.Context, tx *sql.Tx, c entity.Category) entity.Category {
	query := "insert into categories(name) values(?)"
	result, err := tx.ExecContext(ctx, query, c.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	c.Id = int(id)
	return c
}

func (r *CategoryRepository) Update(ctx context.Context, tx *sql.Tx, c entity.Category) entity.Category {
	query := "update categories set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, query, c.Name, c.Id)
	helper.PanicIfError(err)

	return c
}

func (r *CategoryRepository) Delete(ctx context.Context, tx *sql.Tx, c entity.Category) {
	query := "delete from categories where id = ?"
	_, err := tx.ExecContext(ctx, query, c.Id)
	helper.PanicIfError(err)
}

func (r *CategoryRepository) Get(ctx context.Context, tx *sql.Tx) []entity.Category {
	query := "select id, name from categories"
	result, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)

	categories := []entity.Category{}
	for result.Next() {
		category := entity.Category{}
		err := result.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}

func (r *CategoryRepository) GetById(ctx context.Context, tx *sql.Tx, id int) (entity.Category, error) {
	query := "select id, name from categories where id = ?"
	result, err := tx.QueryContext(ctx, query, id)
	helper.PanicIfError(err)

	category := entity.Category{}
	if result.Next() {
		err := result.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		return category, nil
	} else {
		return category, errors.New("cateogry not found")
	}

}
