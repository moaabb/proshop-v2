package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/moaabb/ecommerce/user_svc/infra/config"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg *config.Config, l *zap.Logger) *sql.DB {
	db, err := gorm.Open(postgres.Open(cfg.Dsn), &gorm.Config{})
	if err != nil {
		l.Fatal(fmt.Sprintf("could not connect to db %v", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		l.Fatal(fmt.Sprintf("could not connect to db %v", err))
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
		l.Fatal(fmt.Sprintf("failed to establish connection to db %v", err))
	}

	return sqlDB
}
