package database

import (
	"backend/config"
	"backend/internal/domain/url"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	DB *gorm.DB
}

var dbInstance *Database
var once sync.Once

func ConnectDatabase() *Database {
	once.Do(func() {
		dsn := fmt.Sprintf("postgresql://%s:%s@%s:5432/%s?sslmode=disable", config.AppConfig.Database.Username, config.AppConfig.Database.Password, config.AppConfig.Database.Host, config.AppConfig.Database.Database)
		fmt.Println("DSN : ", dsn)

		//define database logger
		dbLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			},
		)

		dbConfig := gorm.Config{
			Logger: dbLogger,
		}
		//TODO : Change this driver if you use beside postgres (ex: mysql.Open())
		db, err := gorm.Open(postgres.Open(dsn), &dbConfig)
		if err != nil {
			panic("Failed to connect to the database")
		}

		err = db.AutoMigrate(
			url.Url{},
		)
		if err != nil {
			panic("Failed to connect to the database")
		}
		dbInstance = &Database{DB: db}
	})

	return dbInstance
}
