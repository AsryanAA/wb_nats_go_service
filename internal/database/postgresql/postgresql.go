package postgresql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"wb_nats_go_service/internal/config"
	"wb_nats_go_service/internal/models"
)

func Init(cfg *config.Config) *gorm.DB {
	connectPath := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.User, cfg.Password, cfg.DataBaseConfig.Host, cfg.DataBaseConfig.Port, cfg.DBName)
	db, err := gorm.Open(postgres.Open(connectPath), &gorm.Config{})

	if err != nil {
		log.Fatalf("Cannot connect to database, %s", err)
	}

	err = db.AutoMigrate(&models.Item{})
	err = db.AutoMigrate(&models.OrderForDB{})
	if err != nil {
		log.Fatalf("Cannot migrate Order, %s", err)
	}

	return db
}
