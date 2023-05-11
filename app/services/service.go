package service

import (
	"go.uber.org/dig"
)

func Inject(container *dig.Container) error {
	_ = container.Provide(NewUserService)
	_ = container.Provide(NewAuthService)
	_ = container.Provide(NewJWTService)
	_ = container.Provide(NewCategoryService)
	_ = container.Provide(NewCommentService)
	_ = container.Provide(NewOrderService)
	_ = container.Provide(NewProductService)
	_ = container.Provide(NewRatingService)
	return nil
}