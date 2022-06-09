package config

import (
	"net/http"

	"github.com/JoaoGabrielManochio/webapi-go/database"
	"github.com/JoaoGabrielManochio/webapi-go/repository/transaction"
	"github.com/JoaoGabrielManochio/webapi-go/repository/user"
	"github.com/JoaoGabrielManochio/webapi-go/repository/wallet"
	transaction_service "github.com/JoaoGabrielManochio/webapi-go/service/transaction"
	"gorm.io/gorm"
)

var (
	apiDb                 *gorm.DB
	UserRepository        user.IUserRepository
	WalletRepository      wallet.IWalletRepository
	TransactionRepository transaction.ITransactionRepository
	TransactionService    transaction_service.ITransactionService
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
	TransactionRepository = transaction.NewTransaction(apiDb)
	TransactionService = transaction_service.NewService(&http.Client{})

	return err
}
