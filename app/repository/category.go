package repository

import (
	"github.com/dilyara4949/pq_daq/app/models"
	"github.com/dilyara4949/pq_daq/db"
)

type CategoryRepository interface {
	GetById(ID uint) (*models.Category, error)
	GetAll() ([]*models.Category, error) 
	DeleteCategory(category *models.Category) (*models.Category, error) 
}

type CategoryRepo struct {
	Db *db.Database
}

func NewCategoryRepo() CategoryRepository {
	return &CategoryRepo{Db: &db.DB}
}

func (c *CategoryRepo) GetById(ID uint) (*models.Category, error) {
	category := models.Category{}

	if result := c.Db.Db.Select("id", "name").Where("id = ?", ID).First(&category); result.Error != nil {
		return nil, result.Error
	}

	return &category, nil
}

func (c *CategoryRepo) GetAll() ([]*models.Category, error) {
	var categories []*models.Category

	if result := c.Db.Db.Select("id", "name").Find(categories); result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}

func (c *CategoryRepo) DeleteCategory(category *models.Category) (*models.Category, error) {
	if err := c.Db.Db.Delete(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}