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
	GetCategorytByID(ctx context.Context, id uint) (*models.Category, error)
	CreateRating(ctx context.Context, rating *models.Rating) (*models.Rating, error)
	SortByPrice(ctx context.Context, sort string)([]*models.Product, error) 
	SortByRating(ctx context.Context, sort string)([]*models.Product, error) 
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

func(p *productService) GetCategorytByID(ctx context.Context, id uint) (*models.Category, error) {
	category, err := p.product.GetCategortById( id)
	if err != nil {
		return nil, err
	}

	return category, nil
}
func(p *productService) GetAll(ctx context.Context)([]*models.Product, error) {
	products, err := p.product.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}
func(p *productService) SortByPrice(ctx context.Context, sort string)([]*models.Product, error) {
	products, err := p.product.SortByPrice(sort)
	if err != nil {
		return nil, err
	}
	return products, nil
}
func(p *productService) SortByRating(ctx context.Context, sort string)([]*models.Product, error) {
	products, err := p.product.SortByRating(sort)
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

func(p *productService) CreateRating(ctx context.Context, rating *models.Rating) (*models.Rating, error) {
	rating, err := p.product.CreateRating(rating)
	if err != nil {
		return nil, err
	}

	return rating, nil
}