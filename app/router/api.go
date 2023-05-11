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
		rating *c.RatingController,
		comment *c.CommentController,
		order *c.OrderController,
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
		apiCategory.Use(middleware.AuthorizeJWT(jwtS))
		apiCategory.GET("/category", category.GetAll)
		apiCategory.GET("/category/:id", category.GetById)
		apiCategory.POST("/category", category.Create)
		}
		apiProduct := r.Group("api")
		{
		apiCategory.Use(middleware.AuthorizeJWT(jwtS))
		apiProduct.GET("/product", product.GetAll)
		apiProduct.GET("/product/price", product.SortByPrice)
		apiProduct.GET("/product/rating", product.SortByRating)
		apiProduct.GET("/product/:id", product.GetById)
		apiProduct.POST("/product", product.CreateProduct)
		}
		apiComment := r.Group("api")
		{
		apiCategory.Use(middleware.AuthorizeJWT(jwtS))
		apiComment.POST("/comment", comment.AddComment)
		apiComment.GET("/comment", comment.GetAll)
		apiComment.GET("/comment/:product_id", comment.GetByProduct)
		}
		apiOrder := r.Group("api")
		{
			// apiOrder.Use(middleware.AuthorizeJWT(jwtS))
			apiOrder.GET("/order", order.GetAllOrders)
			apiOrder.GET("/order/product/:product_id", order.GetByProduct)
			apiOrder.GET("/order/user/:user_id", order.GetByUser)
			apiOrder.PUT("/order/:id/pay", order.Pay)
			apiOrder.GET("/order/:product_id", order.GetByProduct)
			apiOrder.POST("/order", order.CreateOrder)
		}
		apiRating := r.Group("api")
		{
			apiRating.Use(middleware.AuthorizeJWT(jwtS))
			apiRating.GET("/rating/:product_id", rating.GetById)
			apiRating.PUT("/rating/:product_id", rating.Create)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}
