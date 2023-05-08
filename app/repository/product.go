package repository

import (
	"github.com/dilyara4949/pq_daq/app/models"
	"github.com/dilyara4949/pq_daq/db"
)

type Product struct {
	Db *db.Database
}

func NewProduct() *Product {
	return &Product{Db: &db.DB}
}

func (s *Product)GetById(ID int) (models.Product, error) {

	product := models.Product{}

	if result := s.Db.Db.Select("id", "name", "price", "description").Where("id = ?", ID).First(&product); result.Error != nil {
		return product, result.Error
	}

	return product, nil

}

func (s *Product)GetByName(name string) (models.Product, error) {

	product := models.Product{}

	if result := s.Db.Db.Select("id", "name", "price", "description").Where("name = ?", name).Find(&product); result.Error != nil {
		return product, result.Error
	}

	return product, nil

}

func (s *Product)GetAllProducts() ([]models.Product, error) {
	var products []models.Product

	if result := s.Db.Db.Select("id", "name", "price", "description").Find(&products); result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (s *Product)CreateProduct(product *models.Product) (*models.Product, error) {
	if err := s.Db.Db.Create(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (s *Product)UpdateProduct(product *models.Product) (*models.Product, error) {
	if err := s.Db.Db.Save(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (s *Product)DeleteProduct(product *models.Product) (*models.Product, error) {
	if err := s.Db.Db.Delete(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}