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
	GetCategortById(ID uint) (*models.Category, error) 
}

type ProductRepo struct {
	Db *db.Database
}

func NewProductRepo() ProductRepository {
	return &ProductRepo{Db: &db.DB}
}

func (p *ProductRepo) GetById(ID uint) (*models.Product, error) {
	product := models.Product{}

	if result := p.Db.Db.Where("id = ?", ID).First(&product); result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}


func (p *ProductRepo) GetCategortById(ID uint) (*models.Category, error) {
	category := models.Category{}
	
	if result := p.Db.Db.First(&category, ID); result.Error != nil {
		return &category, result.Error
	}

	return &category, nil
}
func (p *ProductRepo)GetAll() ([]*models.Product, error) {
	var products []*models.Product

	if result := p.Db.Db.Find(&products); result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (p *ProductRepo)CreateProduct(product2 *models.Product) (*models.Product, error) {

	if err := p.Db.Db.Create(&product2).Error; err != nil {
		return nil, err
	}

	return product2, nil
}

func (p *ProductRepo) DeleteProduct(product *models.Product) (*models.Product, error) {
	if err := p.Db.Db.Delete(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}