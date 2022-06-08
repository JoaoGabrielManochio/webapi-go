package wallet

import (
	"errors"
	"net/http"
	"time"

	"github.com/JoaoGabrielManochio/webapi-go/models"
	"github.com/JoaoGabrielManochio/webapi-go/repository/wallet"
	"gorm.io/gorm"
)

// IWalletBusiness : interface of user business
type IWalletBusiness interface {
	PostWallet(wallet models.Wallet) (int, *models.Wallet, error)
	UpdateWallet(wallet models.Wallet) (int, *models.Wallet, error)
	GetWallet(id int64) (int, *models.Wallet, error)
	GetWallets() (int, *[]models.Wallet, error)
	DeleteWallet(id int64) (int, error)
}

// WalletBusiness : struct of user business
type WalletBusiness struct {
	WalletRepository wallet.IWalletRepository
}

// NewWalletBusiness : create a new user business
func NewWalletBusiness(WalletRepository wallet.IWalletRepository) IWalletBusiness {
	return &WalletBusiness{WalletRepository}
}

// PostUser : post user
func (a *WalletBusiness) PostWallet(wallet models.Wallet) (int, *models.Wallet, error) {

	_, err := a.WalletRepository.GetWalletByUserId(wallet.UserId)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusInternalServerError, nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusBadRequest, nil, errors.New("user not found")
	}

	newWallet, err := a.WalletRepository.Create(&models.Wallet{
		Name:      wallet.Name,
		UserId:    wallet.UserId,
		Value:     wallet.Value,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusAccepted, newWallet, nil
}

// GetUser : get user from database by ID
func (a *WalletBusiness) GetWallet(id int64) (int, *models.Wallet, error) {
	wallet, err := a.WalletRepository.FindById(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusBadRequest, nil, errors.New("wallet not found")
	}

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, wallet, nil
}

// GetUsers : get all users from database
func (a *WalletBusiness) GetWallets() (int, *[]models.Wallet, error) {
	wallet, err := a.WalletRepository.FindAll()

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusBadRequest, nil, errors.New("wallets not found")
	}

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, wallet, nil
}

// UpdateUser : update user infos
func (a *WalletBusiness) UpdateWallet(wallet models.Wallet) (int, *models.Wallet, error) {

	w, err := a.WalletRepository.FindById(int64(wallet.Id))

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusBadRequest, nil, errors.New("wallet not found")
	}

	if w.UserId != wallet.UserId {
		return http.StatusBadRequest, nil, errors.New("user_id don't belong to this wallet")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusInternalServerError, nil, err
	}

	newWallet, err := a.WalletRepository.Update(&models.Wallet{
		Id:        wallet.Id,
		Name:      wallet.Name,
		UserId:    wallet.UserId,
		Value:     wallet.Value,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusAccepted, newWallet, nil
}

// DeleteWallet : delete user from database by ID
func (a *WalletBusiness) DeleteWallet(id int64) (int, error) {

	w, err := a.WalletRepository.FindById(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusBadRequest, errors.New("wallet not found")
	}

	if w.Value > 0 {
		return http.StatusBadRequest, errors.New("wallet canot be deleted with value")
	}

	err = a.WalletRepository.Delete(&models.Wallet{}, id)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
