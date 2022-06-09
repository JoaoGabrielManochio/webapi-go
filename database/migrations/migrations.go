package migrations

import (
	"github.com/JoaoGabrielManochio/webapi-go/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Wallet{}, models.User{}, models.Transaction{})
}
