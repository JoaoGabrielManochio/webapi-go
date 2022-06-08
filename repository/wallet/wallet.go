package wallet

import (
	"github.com/JoaoGabrielManochio/webapi-go/models"
	"gorm.io/gorm"
)

type IWalletRepository interface {
	Create(wallet *models.Wallet) (*models.Wallet, error)
	GetWalletByUserId(user_id int64) (*[]models.User, error)
	Update(wallet *models.Wallet) (*models.Wallet, error)
	Delete(wallet *models.Wallet, id int64) error
	FindById(id int64) (*models.Wallet, error)
	FindAll() (*[]models.Wallet, error)
}

// Wallet : struct of wallet repository
type WalletRepository struct{ db *gorm.DB }

// NewWallet : create a new wallet repository
func NewWallet(db *gorm.DB) IWalletRepository {
	return &WalletRepository{db}
}

// FindById : get wallet by ID
func (b *WalletRepository) FindById(id int64) (*models.Wallet, error) {
	wallet := &models.Wallet{}
	find := b.db

	find = find.Where("id = ?", id)

	err := find.
		First(wallet).
		Error

	return wallet, err
}

// FindAll : get all wallet
func (b *WalletRepository) FindAll() (*[]models.Wallet, error) {
	wallet := &[]models.Wallet{}
	find := b.db

	err := find.
		Find(wallet).
		Error

	return wallet, err
}

// Create : create wallet
func (b *WalletRepository) Create(wallet *models.Wallet) (*models.Wallet, error) {

	create := b.db

	err := create.
		Create(wallet).
		Error

	return wallet, err
}

// GetWalletByUserId : get user by user_id
func (b *WalletRepository) GetWalletByUserId(user_id int64) (*[]models.User, error) {

	user := &[]models.User{}
	find := b.db

	find = find.Where("id = ?", user_id)

	err := find.First(user)

	return user, err.Error
}

// Update : update wallet
func (b *WalletRepository) Update(wallet *models.Wallet) (*models.Wallet, error) {

	create := b.db

	err := create.
		Save(wallet).
		Error

	return wallet, err
}

// Delete : delete wallet by ID
func (b *WalletRepository) Delete(wallet *models.Wallet, id int64) error {

	create := b.db

	err := create.
		Delete(wallet, id).
		Error

	return err
}
