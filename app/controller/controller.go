package controller

import (
	"go.uber.org/dig"
)


func Inject(container *dig.Container) error {
	_ = container.Provide(NewUserController)
	_ = container.Provide(NewAuthController)
	return nil
}