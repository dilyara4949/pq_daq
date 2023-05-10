package repository

import (
	"go.uber.org/dig"
)

func Inject(container *dig.Container) error {
	_ = container.Provide(NewUserRepo)
	_ = container.Provide(NewProductRepo)
	_ = container.Provide(NewOrderRepo)
	_ = container.Provide(NewCommentRepo)
	_ = container.Provide(NewCategoryRepo)
	return nil
}