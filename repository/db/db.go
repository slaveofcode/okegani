package db

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/slaveofcode/okegani/repository/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

type DBConfig struct {
	Host    string
	Port    string
	User    string
	Pass    string
	DBName  string
	SSLMode bool
}

func WithEnvDBConfig() DBConfig {
	sslMode, _ := strconv.ParseBool(os.Getenv("PG_SSLMODE"))
	return DBConfig{
		Host:    os.Getenv("PG_HOST"),
		Port:    os.Getenv("PG_PORT"),
		User:    os.Getenv("PG_USER"),
		Pass:    os.Getenv("PG_PASS"),
		DBName:  os.Getenv("PG_DBNAME"),
		SSLMode: sslMode,
	}
}

func Init(config DBConfig) error {
	ssl := "disable"
	if config.SSLMode {
		ssl = "enable"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", config.Host, config.User, config.Pass, config.DBName, config.Port, ssl)
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// disable transaction that enable by default on any operation
		// see: https://gorm.io/docs/transactions.html#Disable-Default-Transaction
		SkipDefaultTransaction: true,
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Warn,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
	})

	if err != nil {
		panic("Couldn't connect to DB:" + err.Error())
	}

	db = dbConn

	return db.AutoMigrate(
		&models.Account{},
		&models.AccountCredential{},
		&models.Category{},
		&models.Post{},
		&models.PostVote{},
		&models.Discussion{},
	)
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			log.Println("Unable close DB:", err.Error())
			return
		}

		sqlDB.Close()
	}
}
