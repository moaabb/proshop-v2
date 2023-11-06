package database

import (
	"database/sql"
	"time"

	"github.com/moaabb/ecommerce/product_svc/infra/config"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg *config.Config, l *zap.Logger) *sql.DB {
	db, err := gorm.Open(postgres.Open(cfg.Dsn), &gorm.Config{})
	if err != nil {
		l.Fatal("could not connect to db %v", zap.Error(err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		l.Fatal("could not connect to db %v", zap.Error(err))
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Ping
	err = sqlDB.Ping()
	if err != nil {
		l.Fatal("failed to establish connection to db %v", zap.Error(err))
	}

	return sqlDB
}
