package service 

import (
	"context"

	"github.com/dilyara4949/pq_daq/app/models"
	repo "github.com/dilyara4949/pq_daq/app/repository"
)

type OrderService interface {
	GetOrderByID(ctx context.Context, id uint) (*models.Order, error)
	GetAll(ctx context.Context, order *models.Order)([]*models.Order, error) 
	DeleteOrder(ctx context.Context, order *models.Order) (*models.Order, error)
	CreateOrder(ctx context.Context, order *models.Order)(*models.Order, error)
	GetProductByID(ctx context.Context, productId uint) (*models.Product, error)
	GetByUser(ctx context.Context, id uint) (*models.Order, error)
	GetByProduct(ctx context.Context, id uint) (*models.Order, error)
	Update(ctx context.Context, order *models.Order) (*models.Order, error)
}

type orderService struct {
	order repo.OrderRepository
}
func NewOrderService(repo repo.OrderRepository) OrderService {
	return &orderService{order: repo}
}

func (u *orderService) Update(ctx context.Context, order *models.Order) (*models.Order, error) {
	order, err := u.order.Update(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func(o *orderService) GetByUser(ctx context.Context, id uint) (*models.Order, error) {
	order, err := o.order.GetByUser(id)
	if err != nil {
		return nil, err
	}

	return order, nil
}
func(o *orderService) GetByProduct(ctx context.Context, id uint) (*models.Order, error) {
	order, err := o.order.GetByProduct(id)
	if err != nil {
		return nil, err
	}

	return order, nil
}
func(o *orderService) GetOrderByID(ctx context.Context, id uint) (*models.Order, error) {
	order, err := o.order.GetById(id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o *orderService) DeleteOrder(ctx context.Context, order *models.Order) (*models.Order, error) {
	order, err := o.order.DeleteOrder(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o *orderService) GetAll(ctx context.Context, order *models.Order)([]*models.Order, error) {
	orders, err := o.order.GetAll()
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (o *orderService) CreateOrder(ctx context.Context, order *models.Order)(*models.Order, error) {
	order, err := o.order.CreateOrder(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func(p *orderService) GetProductByID(ctx context.Context, productId uint) (*models.Product, error) {
	product, err := p.order.GetByIdProduct(productId)
	if err != nil {
		return nil, err
	}

	return product, nil
}