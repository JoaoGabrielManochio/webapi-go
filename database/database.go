package database

import (
	"fmt"
	"log"

	"github.com/JoaoGabrielManochio/webapi-go/database/migrations"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// StartDB : open MySQL database connection
func StartDB(host string, port int64, user string, dbName string, password string) (*gorm.DB, error) {
	sn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbName)

	db, err := gorm.Open(mysql.Open(sn), &gorm.Config{})
	if err != nil {
		msg := fmt.Sprintf("Database connection error: %s", err)

		log.Fatal("errror: ", msg)

		return nil, err
	}

	return db, nil
}

func GetDatabase() *gorm.DB {
	return db
}

func Load(db *gorm.DB) {
	migrations.RunMigrations(db)
}
