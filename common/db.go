package common

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type DBConfig struct {
	Host     string
	Username string
	Password string
	DBName   string
	Port     int
}

func NewDB() {
	db_config := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     3306,
	}
	conn_string := fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True&loc=Local", db_config.Username, db_config.Password, db_config.DBName)
	db, err = gorm.Open("mysql", conn_string)
	if err != nil {
		logrus.WithError(err).Panic("Error when connecting database")
	}
	logrus.Info("Database Connected", db)
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	err := db.Close()
	if err != nil {
		logrus.WithError(err).Panic("Error when closing database")
	}
	fmt.Println("=== DB Closed ===")
}
