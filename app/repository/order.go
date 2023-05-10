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
}

type OrderRepo struct {
	Db *db.Database
}

func NewOrderRepo() UserRepository {
	return &UserRepo{Db: &db.DB}
}

func (s *OrderRepo) GetById(ID uint) (*models.Order, error) {

	order := models.Order{}

	if result := s.Db.Db.Select("id", "PaymentStatus", "TotalPrice").Where("id=?", ID).First(&order); result.Error != nil {
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