package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/adhyttungga/go-crud-gin-swagger/model/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	Err error
)

func ConnectDB() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// SQL Server
	// dsn := fmt.Sprintf("%s://%s:%s@%s:%s?database=%s", Config.DB.Adapter, Config.DB.User, Config.DB.Password, Config.DB.Host, Config.DB.Port, Config.DB.Name)
	// Postgre SQL
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Config.DB.Host, Config.DB.Port, Config.DB.User, Config.DB.Password, Config.DB.Name)
	DB, Err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if Err != nil {
		panic(Err)
	}

	DB.Table("tags").AutoMigrate(&entity.Tags{})

	return DB
}
