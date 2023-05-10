package repository

import (
	"github.com/dilyara4949/pq_daq/app/models"
	"github.com/dilyara4949/pq_daq/db"
	// "gorm.io/gorm"
)

type UserRepository interface {
	GetById(ID uint) (*models.User, error)
	GetAll() ([]*models.User, error) 
	CreateUser(user *models.User) (*models.User, error) 
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(user *models.User) (*models.User, error) 
}

type UserRepo struct {
	Db *db.Database
}



func NewUserRepo() UserRepository {
	return &UserRepo{Db: &db.DB}
}

func (s *UserRepo)GetById(ID uint) (*models.User, error) {

	user := models.User{}

	if result := s.Db.Db.Select("id", "firstname", "lastname", "email").Where("id = ?", ID).First(&user); result.Error != nil {
		return &user, result.Error
	}

	return &user, nil

}



func (s *UserRepo)GetAll() ([]*models.User, error) {
	var users []*models.User

	if result := s.Db.Db.Find(&users); result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (s *UserRepo)CreateUser(user *models.User) (*models.User, error) {
	if err := s.Db.Db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (s *UserRepo)UpdateUser(user *models.User) (*models.User, error) {
	if err := s.Db.Db.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (s *UserRepo)DeleteUser(user *models.User) (*models.User, error) {
	if err := s.Db.Db.Delete(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}