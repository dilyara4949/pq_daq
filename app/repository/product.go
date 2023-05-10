package repository

import (
	"github.com/dilyara4949/pq_daq/app/models"
	"github.com/dilyara4949/pq_daq/db"
)

type ProductRepository interface {
	GetById(ID uint) (*models.Product, error)
	GetAll() ([]*models.Product, error)
	CreateProduct(product *models.Product) (*models.Product, error)
	DeleteProduct(product *models.Product) (*models.Product, error)
}

type ProductRepo struct {
	Db *db.Database
}

func NewProductRepo() ProductRepository {
	return &ProductRepo{Db: &db.DB}
}

func (p *ProductRepo) GetById(ID uint) (*models.Product, error) {
	product := models.Product{}

	if result := p.Db.Db.Find(&product); result.Error != nil {
		return &product, result.Error
	}

	return &product, nil
}

func (p *ProductRepo)GetAll() ([]*models.Product, error) {
	var products []*models.Product

	if result := p.Db.Db.Find(&products); result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (p *ProductRepo)CreateProduct(product *models.Product) (*models.Product, error) {
	if err := p.Db.Db.Create(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (p *ProductRepo) DeleteProduct(product *models.Product) (*models.Product, error) {
	if err := p.Db.Db.Delete(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}