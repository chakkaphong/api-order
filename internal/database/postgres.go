package database

import (
	"api-order/configs"
	"api-order/internal/entities"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error

	// DB returns the underlying *gorm.DB instance
	DB() *gorm.DB
}

type service struct {
	db *gorm.DB
}

var (
	database   string
	password   string
	username   string
	port       string
	host       string
	ssl        string
	dbInstance *service
	env        string
)

func init() {
	secretCfg := configs.GetSecret()
	database = secretCfg.Postgres.Db
	password = secretCfg.Postgres.Password
	username = secretCfg.Postgres.User
	port = secretCfg.Postgres.Port
	host = secretCfg.Postgres.Host
	ssl = secretCfg.Postgres.Ssl

	cfg := configs.GetConfig()
	env = cfg.App.Env
}

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC+0",
		host, username, password, database, port)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		QueryFields: true,
		Logger:      GormLogger{logger.Default.LogMode(logger.Info)},
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(
		entities.Orders{},
		entities.OrderItems{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Set connection pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

type GormLogger struct {
	logger.Interface
}

func (s *service) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database connection: %v", err)
	}

	return sqlDB.Close()
}

func (s *service) DB() *gorm.DB {
	return s.db
}
