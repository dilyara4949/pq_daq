package repository

import (
	"github.com/dilyara4949/pq_daq/app/models"
	"github.com/dilyara4949/pq_daq/db"
)

type OrderRepository interface {
	GetById(ID uint) (*models.Order, error)
	CreateOrder(order *models.Order) (*models.Order, error)
	DeleteOrder(order *models.Order) (*models.Order, error) 
	GetAll() ([]*models.Order, error)
	GetByIdProduct(ID uint) (*models.Product, error) 
	GetByUser(ID uint) (*models.Order, error)
	GetByProduct(ID uint) (*models.Order, error)
	Update(order *models.Order) (*models.Order, error)
}

type OrderRepo struct {
	Db *db.Database
}

func NewOrderRepo() OrderRepository {
	return &OrderRepo{Db: &db.DB}
}
func (s *OrderRepo)Update(order *models.Order) (*models.Order, error) {
	if err := s.Db.Db.Save(&order).Error; err != nil {
		return order, err
	}

	return order, nil	
}

func (s *OrderRepo) GetById(ID uint) (*models.Order, error) {

	order := models.Order{}

	if result := s.Db.Db.Where("id=?", ID).First(&order); result.Error != nil {
		return &order, result.Error
	}

	return &order, nil 
}

func (s *OrderRepo) GetByUser(ID uint) (*models.Order, error) {

	order := models.Order{}

	if result := s.Db.Db.Where("user_id=?", ID).First(&order); result.Error != nil {
		return &order, result.Error
	}

	return &order, nil 
}
func (s *OrderRepo) GetByProduct(ID uint) (*models.Order, error) {

	order := models.Order{}

	if result := s.Db.Db.Where("product_id=?", ID).First(&order); result.Error != nil {
		return &order, result.Error
	}

	return &order, nil 
}

func (s *OrderRepo)CreateOrder(order *models.Order) (*models.Order, error) {

	if err := s.Db.Db.Create(&order).Error; err != nil {
		return order, err
	}

	return order, nil
}

func (s *OrderRepo)DeleteOrder(order *models.Order) (*models.Order, error) {

	if err := s.Db.Db.Delete(&order).Error; err != nil {
		return order, err
	}

	return order, nil
}

func (s *OrderRepo)GetAll() ([]*models.Order, error) {
	var orders []*models.Order 

	if result := s.Db.Db.Select("id" , "PaymentStatus", "TotalPrice").Find(&orders); result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

func (p *OrderRepo) GetByIdProduct(ID uint) (*models.Product, error) {
	product := models.Product{}

	if result := p.Db.Db.Where("id = ?", ID).First(&product); result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}