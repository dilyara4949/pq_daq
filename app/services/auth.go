package service

import (
	"log"
	"github.com/dilyara4949/pq_daq/app/models"
	"github.com/dilyara4949/pq_daq/app/repository"
	"golang.org/x/crypto/bcrypt"
	"github.com/mashingan/smapping"

)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	// VerifyCredential(email string, password string) *models.User
	CreateUser(user models.User) *models.User
	// FindByEmail(email string) models.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{userRepository: userRep}
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.userRepository.VerifyCredential(email, password)
	if v, ok := res.(models.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}
// func (service *authService) VerifyCredential(email string, password string) interface{} {
// 	res := service.userRepository.VerifyCredential(email, password)
// 	if res != nil {
// 		user := res.(*models.User)
// 		comparedPassword := comparePassword(user.Password, []byte(password))
// 		if user.Email == email && comparedPassword {
// 			return user
// 		}
// 	}
// 	return nil
// }


// func (service *authService) CreateUser(user models.Register) *models.User {
// 	userToCreate := models.User{}
// 	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
// 	if err != nil {
// 		log.Fatalf("Error while creating %v", err)
// 	}
// 	res, err := service.userRepository.CreateUser(&userToCreate)
// 	return res
// }
func (service *authService) CreateUser(user models.User) *models.User {
	userToCreate := models.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Error while creating %v", err)
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userToCreate.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error while hashing password %v", err)
	}

	userToCreate.Password = string(hashedPassword)

	res, err := service.userRepository.CreateUser(&userToCreate)
	return res
}
func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

// func comparePassword(hashedPwd string, plainPassword []byte) bool {
// 	byteHash := []byte(hashedPwd)
// 	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
// 	if err != nil {
// 		log.Println(err)
// 		return false
// 	}
// 	return true
// }
func comparePassword(hashedPassword string, plainPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}