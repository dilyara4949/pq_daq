package controller

import (
	"go.uber.org/dig"
)


func Inject(container *dig.Container) error {
	_ = container.Provide(NewUserAPI)
	return nil
}