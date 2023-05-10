package router

import (
	"log"

	c "github.com/dilyara4949/pq_daq/app/controller"
	"github.com/dilyara4949/pq_daq/app/middleware"
	s "github.com/dilyara4949/pq_daq/app/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func RunAPI(r *gin.Engine, container *dig.Container) error {
	err := container.Invoke(func(
		user *c.UserController,
		auth *c.AuthController,
		product *c.ProductController,
		category *c.CategoryController,
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
		apiCategory := r.Group("api")
		{
		// apiCategory.Use(middleware.AuthorizeJWT(jwtS))
		apiCategory.GET("/category", category.GetAll)
		apiCategory.GET("/category/:id", category.GetById)
		apiCategory.POST("/category", category.Create)
		}
		apiProduct := r.Group("api")
		{
		// apiCategory.Use(middleware.AuthorizeJWT(jwtS))
		apiProduct.GET("/product", product.GetAll)
		apiProduct.GET("/product/:id", product.GetById)
		apiProduct.POST("/product", product.CreateProduct)
		}
		// apiComment := r.Group("api")
		// {
		// 	apiComment.Use(middleware.AuthorizeJWT(jwtS))
		// 	apiComment.GET("/comment/:priduct_id", comment.GetAll)
		// 	apiComment.POST("/comment:priduct_id", comment.Create)
		// }
		// apiOrder := r.Group("api")
		// {
		// 	apiOrder.Use(middleware.AuthorizeJWT(jwtS))
		// 	apiOrder.GET("/order", order.GetAll)
		// 	apiOrder.GET("/order/:priduct_id", order.GetById)
		// 	apiOrder.POST("/order/:priduct_id", order.Create)
		// }
		// apiRating := r.Group("api")
		// {
		// 	apiRating.Use(middleware.AuthorizeJWT(jwtS))
		// 	apiRating.GET("/rating/:priduct_id", rating.GetById)
		// 	apiRating.PUT("/rating/:priduct_id", rating.Create)
		// }

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}
