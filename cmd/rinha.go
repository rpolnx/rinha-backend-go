package main

import (
	"fmt"
	"log"

	configs "github.com/rpolnx/rinha-backend-go/internal/config"
	"github.com/rpolnx/rinha-backend-go/internal/controller"
	"github.com/rpolnx/rinha-backend-go/internal/repository"
	"github.com/rpolnx/rinha-backend-go/internal/routes"
	"github.com/rpolnx/rinha-backend-go/internal/server"
	"github.com/rpolnx/rinha-backend-go/internal/service"

	"github.com/caarlos0/env/v10"
)

func main() {

	cfg := configs.EnvConfig{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("error serializing envs: %v", err)
	}

	db, err := repository.InitializeDb(cfg)

	if err != nil {
		log.Fatalf("error initializing db: %v", err)
	}

	personRepo := repository.NewPersonRepository(db)
	personSvc := service.NewPersonService(personRepo)
	personCtrl := controller.NewPersonController(personSvc)

	m := server.NewServer()

	routes.
		NewRouterBuilder(m, personCtrl).
		AppendRoutes()

	if err := m.Run(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		log.Fatalf("exiting program with error: %v", err)
	}
}
