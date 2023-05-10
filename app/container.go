package app

import (
	"log"

	"go.uber.org/dig"

	"github.com/dilyara4949/pq_daq/app/controller"
	"github.com/dilyara4949/pq_daq/app/repository"
	"github.com/dilyara4949/pq_daq/app/services"
)

func Build() *dig.Container {
	container := dig.New()

	err := repository.Inject(container)
	if err != nil {
		log.Fatal("Error to inject repositories", err)
	}

	err = service.Inject(container)
	if err != nil {
		log.Fatal("Error to inject services", err)
	}

	err = controller.Inject(container)
	if err != nil {
		log.Fatal("Error to inject APIs", err)
	}

	return container
}