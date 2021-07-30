package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

//Model is sample of common table structure
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")
	postFix := "?charset=utf8mb4,utf8&sql_mode=TRADITIONAL&multiStatements=true&parseTime=true"

	msql := mysql.Config{}
	log.Println(msql)

	var logLevel logger.LogLevel
	if os.Getenv("DB_LOG_ENABLED") == "true" {
		logLevel = logger.Info
	} else {
		logLevel = logger.Silent
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logLevel,
			Colorful:      true,
		},
	)

	conn, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s", username, password, dbHost, dbPort, dbName, postFix)), &gorm.Config{
		Logger:      newLogger,
		PrepareStmt: true,
	})

	if err != nil {
		fmt.Print(err)
	}
	db = conn

	//Automatically create migration as per model
	db.Debug().AutoMigrate(
		&User{},
	)
}

//GetDB function return the instance of db
func GetDB() *gorm.DB {
	return db
}
