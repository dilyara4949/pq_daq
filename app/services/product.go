package service 

import (
	"context"

	"github.com/dilyara4949/pq_daq/app/models"
	repo "github.com/dilyara4949/pq_daq/app/repository"
)

type ProductService interface {
	GetProductByID(ctx context.Context, id uint) (*models.Product, error)
	GetAll(ctx context.Context)([]*models.Product, error)
	DeleteProduct(ctx context.Context, product *models.Product)(*models.Product, error)
	CreateProduct(ctx context.Context, product *models.Product) (*models.Product, error)
}

type productService struct {
	product repo.ProductRepository
}
func NewProductService(repo repo.ProductRepository) ProductService {
	return &productService{product: repo}
}

func(p *productService) GetProductByID(ctx context.Context, id uint) (*models.Product, error) {
	product, err := p.product.GetById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func(p *productService) GetAll(ctx context.Context)([]*models.Product, error) {
	products, err := p.product.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func(p *productService) DeleteProduct(ctx context.Context, product *models.Product)(*models.Product, error) {
	product, err := p.product.DeleteProduct(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func(p *productService) CreateProduct(ctx context.Context, item *models.Product) (*models.Product, error) {
	product, err := p.product.CreateProduct(item)
	if err != nil {
		return nil, err
	}

	return product, nil
}