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
func StartDB(port int64, host, user, dbName, password, timeZone string) (*gorm.DB, error) {
	sn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=%s", user, password, host, port, dbName, timeZone)

	db, err := gorm.Open(mysql.Open(sn), &gorm.Config{})
	if err != nil {
		msg := fmt.Sprintf("Database connection error: %s", err)

		log.Fatal("errror: ", msg)

		return nil, err
	}

	return db, nil
}

func Load(db *gorm.DB) {
	migrations.RunMigrations(db)
}
