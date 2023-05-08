package app

import (
	"log"

	"github.com/dilyara4949/pq_daq/app/repository"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	// Inject repositories
	err := repository.Inject(container)
	if err != nil {
		log.Fatal("Failed to inject repositories", err)
	}
	return container
}