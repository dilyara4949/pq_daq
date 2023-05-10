package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/dilyara4949/pq_daq/app/models"
	service "github.com/dilyara4949/pq_daq/app/services"
	"github.com/gin-gonic/gin"
)

// type AuthController interface {
// 	Login(ctx *gin.Context)
// 	Register(ctx *gin.Context)
// }

type AuthController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) *AuthController {
	return &AuthController{ authService: authService, jwtService:  jwtService}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var login models.Login
	err := ctx.ShouldBind(&login)
	if err != nil {
		response := BuildErrorResponse("Failed to process the request", err.Error(), EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(login.Email, login.Password)
	if v, ok := authResult.(models.User); ok {
		generateToken := c.jwtService.GenerateToken(strconv.FormatUint(uint64(v.ID), 10))
		v.Token = generateToken
		response := BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return 
	}
	response := BuildErrorResponse("Please verify your credentials", "The credentials are incorrect.", EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *AuthController) Register(ctx *gin.Context) {
	// var register models.Register
	var user models.User
	err := ctx.ShouldBind(&user)
	if err != nil {
		response := BuildErrorResponse("Failed to process the request", err.Error(), EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !c.authService.IsDuplicateEmail(user.Email) {
		response := BuildErrorResponse("Failed to process the request", err.Error(), EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createUser := c.authService.CreateUser(user)
		token := c.jwtService.GenerateToken(strconv.FormatUint(uint64(createUser.ID), 10))
		createUser.Token = token
		response := BuildResponse(true, "OK!", createUser)
		ctx.JSON(http.StatusCreated, response)
	}
}



type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Error:   nil,
		Data:    data,
	}
	return res
}

func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Error:   splittedError,
		Data:    data,
	}
	return res
}

type EmptyObj struct{}