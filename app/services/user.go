package service

import (
	"context"

	"github.com/dilyara4949/pq_daq/app/models"
	repo "github.com/dilyara4949/pq_daq/app/repository"
)
type UserService interface {
	GetUserByID(ctx context.Context, id uint) (*models.User, error)
	GetAll(ctx context.Context) ([]*models.User, error)
	DeleteUser(ctx context.Context, user *models.User) (*models.User, error) 
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
}


type userService struct{
	user repo.UserRepository
} 
func NewUserService(repo repo.UserRepository) UserService {
	return &userService{user: repo}
}

func (u *userService) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	user, err := u.user.GetById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) GetAll(ctx context.Context) ([]*models.User, error) {
	users, err := u.user.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userService) DeleteUser(ctx context.Context, user *models.User) (*models.User, error) {
	user, err := u.user.DeleteUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	user, err := u.user.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (u *userService) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	user, err := u.user.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
