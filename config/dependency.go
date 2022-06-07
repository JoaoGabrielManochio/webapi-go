package dependency

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
	apiDb, err := database.StartDB("localhost", 3308, "root", "api", "root")

	if err != nil {
		return err
	}

	database.Load(apiDb)

	UserRepository = user.NewUser(apiDb)

	return err
}
