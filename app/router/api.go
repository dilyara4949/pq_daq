package router

import (
	"log"

	c "github.com/dilyara4949/pq_daq/app/controller"
	s "github.com/dilyara4949/pq_daq/app/services"
	"github.com/dilyara4949/pq_daq/app/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)


func RunAPI(r *gin.Engine, container *dig.Container) error {
	err := container.Invoke(func(
		user *c.UserController,
		auth *c.AuthController,
		jwtS s.JWTService,
	) error {

		//--------------------------------API-----------------------------------
		apiAuth := r.Group("api")
		{
			apiAuth.POST("/login", auth.Login)
			apiAuth.POST("/register", auth.Register)
		}

		apiUser := r.Group("api")
		{
			apiUser.Use(middleware.AuthorizeJWT(jwtS))
			apiUser.GET("/users", user.GetAll)
		}


		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}