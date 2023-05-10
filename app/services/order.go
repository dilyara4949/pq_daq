package service 

import (
	"context"

	"github.com/dilyara4949/pq_daq/app/models"
	repo "github.com/dilyara4949/pq_daq/app/repository"
)

type orderService struct {
	order repo.OrderRepo
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

