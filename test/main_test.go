//go:build integration
// +build integration

package main_test

import (
	"log"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/caarlos0/env/v10"
	configs "github.com/rpolnx/rinha-backend-go/internal/config"
	"github.com/rpolnx/rinha-backend-go/internal/controller"
	"github.com/rpolnx/rinha-backend-go/internal/repository"
	"github.com/rpolnx/rinha-backend-go/internal/server"
	"github.com/rpolnx/rinha-backend-go/internal/service"
)

func TestMain(m *testing.M) {
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

	muxServer := server.NewServer()
	personCtrl.RegisterRoutes(muxServer)

	ts := httptest.NewServer(muxServer)

	defer ts.Close()

	res := m.Run()

	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		log.Fatal(err)
	}

	err = p.Signal(os.Interrupt)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(200 * time.Millisecond)

	os.Exit(res)
}
