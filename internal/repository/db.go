package repository

import (
	"fmt"

	configs "github.com/rpolnx/rinha-backend-go/internal/config"
	"github.com/samber/do"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitializeDb(injector *do.Injector) (*gorm.DB, error) {
	cfg := do.MustInvoke[*configs.EnvConfig](injector)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.DbHost,
		cfg.DbUser,
		cfg.DbPassword,
		cfg.DbDbname,
		cfg.DbPort,
		cfg.DbSslmode,
		cfg.DbTimezone,
	)

	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
