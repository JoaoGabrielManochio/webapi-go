package transaction

import (
	"github.com/JoaoGabrielManochio/webapi-go/models"
	"gorm.io/gorm"
)

type ITransactionRepository interface {
	Create(transaction *models.Transaction) (*models.Transaction, error)
	FindById(id int64) (*[]models.Transaction, error)
	FindAll() (*[]models.Transaction, error)
	GetUserById(id int64) (*models.User, error)
	GetUserWalletByUserId(id uint) (*models.Wallet, error)
	SubtractValueFromWallet(wallet *models.Wallet, valueSubtract float64) error
	AddValueToWallet(wallet *models.Wallet, valueSubtract float64) error
}

// Transaction : struct of Transaction repository
type TransactionRepository struct{ db *gorm.DB }

// NewTransaction : create a new Transaction repository
func NewTransaction(db *gorm.DB) ITransactionRepository {
	return &TransactionRepository{db}
}

// FindById : get transaction by ID
func (b *TransactionRepository) FindById(id int64) (*[]models.Transaction, error) {
	transaction := &[]models.Transaction{}
	find := b.db

	if id != 0 {
		find = find.Where("id = ?", id)
	}

	err := find.
		First(transaction).
		Error

	return transaction, err
}

// FindAll : get all transactions
func (b *TransactionRepository) FindAll() (*[]models.Transaction, error) {
	transaction := &[]models.Transaction{}
	find := b.db

	err := find.
		Find(transaction).
		Error

	return transaction, err
}

// Create : create transaction
func (b *TransactionRepository) Create(transaction *models.Transaction) (*models.Transaction, error) {

	create := b.db

	err := create.
		Create(transaction).
		Error

	return transaction, err
}

// GetUserWalletByUserId : get user wallet by user_id
func (b *TransactionRepository) GetUserWalletByUserId(user_id uint) (*models.Wallet, error) {
	userWallet := &models.Wallet{}
	find := b.db

	find = find.Where("user_id = ?", user_id)

	err := find.First(userWallet)

	return userWallet, err.Error
}

// GetUserById : get user by id
func (b *TransactionRepository) GetUserById(id int64) (*models.User, error) {

	user := &models.User{}
	find := b.db

	find = find.Where("id = ?", id)

	err := find.First(user)

	return user, err.Error
}

// AddValueToWallet : add value to wallet
func (b *TransactionRepository) AddValueToWallet(wallet *models.Wallet, valueSubtract float64) error {

	subtract := b.db

	newValue := wallet.Value + valueSubtract

	err := subtract.
		Model(wallet).
		Update("value", newValue).
		Error

	return err
}

// SubtractValueFromWallet : subtract value from wallet
func (b *TransactionRepository) SubtractValueFromWallet(wallet *models.Wallet, valueSubtract float64) error {

	subtract := b.db

	newValue := wallet.Value - valueSubtract

	err := subtract.
		Model(wallet).
		Update("value", newValue).
		Error

	return err
}
