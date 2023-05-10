package service

import (
	"go.uber.org/dig"
)

func Inject(container *dig.Container) error {
	_ = container.Provide(NewUserService)
	_ = container.Provide(NewAuthService)
	_ = container.Provide(NewJWTService)
	return nil
}