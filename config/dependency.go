package config

import (
	"github.com/JoaoGabrielManochio/webapi-go/database"
	"github.com/JoaoGabrielManochio/webapi-go/repository/user"
	"github.com/JoaoGabrielManochio/webapi-go/repository/wallet"
	"gorm.io/gorm"
)

var (
	apiDb            *gorm.DB
	UserRepository   user.IUserRepository
	WalletRepository wallet.IWalletRepository
)

func Load() error {
	// DB
	apiDb, err := database.StartDB(3308, "localhost", "root", "api", "root", "America%2FSao_Paulo")

	if err != nil {
		return err
	}

	database.Load(apiDb)

	UserRepository = user.NewUser(apiDb)
	WalletRepository = wallet.NewWallet(apiDb)

	return err
}
