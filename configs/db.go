package configs

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var err error

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
	TimeZone string
}

func getConfig() DBConfig {
	return DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     portConv(os.Getenv("DB_PORT")),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		TimeZone: "Asia/Bangkok",
	}
}

// convert string to int
func portConv(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("Failed to convert port to integer: %v", err))
	}
	return i
}

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	// sql, _ := fc()
	// fmt.Printf("%v\n==============================\n", sql)
}

func DatabaseInit() {
	cfg := getConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d TimeZone=Asia/Bangkok",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port)
	
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: &SqlLogger{}, // check sql with log
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

}

func GetDB() *gorm.DB {
	return DB
}
