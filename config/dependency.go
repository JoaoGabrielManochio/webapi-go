package config

import (
	"github.com/JoaoGabrielManochio/webapi-go/database"
	"github.com/JoaoGabrielManochio/webapi-go/repository/user"
	"gorm.io/gorm"
)

var (
	apiDb          *gorm.DB
	UserRepository user.IUserRepository
)

func Load() error {
	// DB
	apiDb, err := database.StartDB(3308, "localhost", "root", "api", "root", "America%2FSao_Paulo")

	if err != nil {
		return err
	}

	database.Load(apiDb)

	UserRepository = user.NewUser(apiDb)

	return err
}
