package service

import (
	"context"

	"github.com/dilyara4949/pq_daq/app/models"
	repo "github.com/dilyara4949/pq_daq/app/repository"
)

type userService struct{
	user repo.UserRepo
} 

func (u *userService) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	user, err := u.user.GetById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
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
