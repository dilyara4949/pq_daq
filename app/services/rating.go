package service

import (
	"context"
	"crypto/rsa"

	"github.com/dilyara4949/pq_daq/app/models"
	repo "github.com/dilyara4949/pq_daq/app/repository"
)

type RatingService interface {
	GetRatingByID(ctx context.Context, id uint) (*models.Rating, error) 
	GetAll(ctx context.Context)([]*models.Rating, error)
	CreateRating(ctx context.Context, rating *models.Rating) (*models.Rating, error) 
	Delete(ctx context.Context, rating *models.Rating)(*models.Rating, error) 
}

type ratingService struct {
	rating repo.RatingRepository
}
func NewRatingService(repo repo.RatingRepository) RatingService {
	return &ratingService{rating: repo}
}

func(p *ratingService) GetRatingByID(ctx context.Context, id uint) (*models.Rating, error) {
	rating, err := p.rating.GetById(id)
	if err != nil {
		return nil, err
	}

	return rating, nil
}


func(p *ratingService) GetAll(ctx context.Context)([]*models.Rating, error) {
	ratings, err := p.rating.GetAll()
	if err != nil {
		return nil, err
	}
	return ratings, nil
}

func(p *ratingService) Delete(ctx context.Context, rating *models.Rating)(*models.Rating, error) {
	rating, err := p.rating.Delete(rating)
	if err != nil {
		return nil, err
	}

	return rating, nil
}

func(p *ratingService) CreateRating(ctx context.Context, rating *models.Rating) (*models.Rating, error) {
	rating, err := p.rating.Create(rating)
	if err != nil {
		return nil, err
	}

	return rating, nil
}

