package service

import (
	"context"

	"github.com/dilyara4949/pq_daq/app/models"
    repo "github.com/dilyara4949/pq_daq/app/repository"
)

type CategoryService interface {
	GetCategoryByID(ctx context.Context, id uint) (*models.Category, error)
	GetAll(ctx context.Context, category []*models.Category) ([]*models.Category, error)
	DeleteCategory(ctx context.Context, category *models.Category) (*models.Category, error) 
}


type categoryService struct{
	category repo.CategoryRepository
}

func NewCategoryService(repo repo.CategoryRepository) CategoryService {
	return &categoryService{category: repo}
}


func (c *categoryService) GetCategoryByID(ctx context.Context, id uint) (*models.Category, error) {
	category, err := c.category.GetById(id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (c *categoryService) DeleteCategory(ctx context.Context, category *models.Category) (*models.Category, error) {
	category, err := c.category.DeleteCategory(category)
	if err!= nil {
		return nil, err
	}

	return category, nil
}

func (c *categoryService) GetAll(ctx context.Context, category []*models.Category) ([]*models.Category, error) {
	categories, err := c.category.GetAll()
	if err != nil {
		return nil, err
	}

	return categories, nil
}
