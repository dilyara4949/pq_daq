package router

import (
	"log"

	c "github.com/dilyara4949/pq_daq/app/controller"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)



func RunAPI(r *gin.Engine, container *dig.Container) error {
	err := container.Invoke(func(
		user *c.UserController,
	) error {

		//--------------------------------API-----------------------------------
		apiV := r.Group("api")
		{
			apiV.GET("/users", user.GetAll)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}