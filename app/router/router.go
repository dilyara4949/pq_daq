package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func GinInit(container *dig.Container) *gin.Engine {
	app := gin.New()
	RunAPI(app, container)
	return app
}