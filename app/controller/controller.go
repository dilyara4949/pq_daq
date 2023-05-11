package controller

import (
	"go.uber.org/dig"
)


func Inject(container *dig.Container) error {
	_ = container.Provide(NewUserController)
	_ = container.Provide(NewAuthController)
	_ = container.Provide(NewCategoryController)
	_ = container.Provide(NewProductController)
	_ = container.Provide(NewRatingController)
	return nil
}