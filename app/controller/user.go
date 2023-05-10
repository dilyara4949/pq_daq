package controller

import (
	"net/http"

	"github.com/dilyara4949/pq_daq/app/services"
	"github.com/gin-gonic/gin"
)



type UserController struct {
	user service.UserService
}

func NewUserAPI(service service.UserService) *UserController {
	return &UserController{user: service}
}


func (u *UserController) GetAll(c *gin.Context) {
	users, err := u.user.GetAll(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, users)
}