package service

import (
	"context"

	"github.com/dilyara4949/pq_daq/app/models"
    repo "github.com/dilyara4949/pq_daq/app/repository"
)

type categoryService struct{
	category repo.CategoryRepo
}

func (c *categoryService) getCategoryByID(ctx context.Context, id uint) (*models.Category, error) {
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

func (c *categoryService) GetAllCategories(ctx context.Context, category *models.Category) ([]models.Category, error) {
	categories, err := c.category.GetAllCategories()
	if err != nil {
		return nil, err
	}

	return categories, nil
}
