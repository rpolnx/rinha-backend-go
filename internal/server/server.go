package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	configs "github.com/rpolnx/rinha-backend-go/internal/config"
	"github.com/rpolnx/rinha-backend-go/internal/controller"
	"github.com/rpolnx/rinha-backend-go/internal/repository"
	route "github.com/rpolnx/rinha-backend-go/internal/routes"
	"github.com/rpolnx/rinha-backend-go/internal/service"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func SetupServer() error {
	injector := do.New()

	do.Provide[*configs.EnvConfig](injector, configs.NewEnvConfig)
	do.Provide[*gorm.DB](injector, repository.InitializeDb)

	do.Provide[repository.PersonRepository](injector, repository.NewPersonRepository)
	do.Provide[service.PersonService](injector, service.NewPersonService)
	do.Provide[controller.PersonController](injector, controller.NewPersonController)

	do.Provide[*gin.Engine](injector, NewServer)

	route.RegisterRoutes(injector)

	return RunServer(injector)
}

func NewServer(injector *do.Injector) (*gin.Engine, error) {

	r := gin.Default()

	return r, nil
}

func RunServer(injector *do.Injector) error {
	cfg := do.MustInvoke[*configs.EnvConfig](injector)
	server := do.MustInvoke[*gin.Engine](injector)

	return server.Run(fmt.Sprintf(":%d", cfg.Port))
}
