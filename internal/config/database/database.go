package database

import (
	"fmt"
	"log"
	"notification-engine/internal/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// OpenConnection cria a conexão e configura o pool
func OpenConnection(cfg *config.Config) (*gorm.DB, error) {
	databaseURL := "postgresql://" + cfg.DBUsername + ":" + cfg.DBPassword + "@" + cfg.DBHost + ":" + cfg.DBPort + "/"

	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar no gorm: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("falha ao obter instancia sql.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("banco de dados nao responde: %w", err)
	}

	log.Println("✅ Conexão GORM estabelecida e Pool configurado!")
	return db, nil
}
