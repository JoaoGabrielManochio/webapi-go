package transaction

import (
	"errors"
	"net/http"
	"time"

	"github.com/JoaoGabrielManochio/webapi-go/config"
	"github.com/JoaoGabrielManochio/webapi-go/models"
	"github.com/JoaoGabrielManochio/webapi-go/repository/transaction"
	transaction_service "github.com/JoaoGabrielManochio/webapi-go/service/transaction"
	"gorm.io/gorm"
)

// ITransactionBusiness : interface of Transaction business
type ITransactionBusiness interface {
	PostTransaction(transaction models.Transaction) (int, *models.Transaction, error)
	GetTransaction(id int64) (int, *[]models.Transaction, error)
	GetTransactions() (int, *[]models.Transaction, error)
}

// TransactionBusiness : struct of Transaction business
type TransactionBusiness struct {
	TransactionRepository transaction.ITransactionRepository
	TransactionService    transaction_service.ITransactionService
}

// NewTransactionBusiness : create a new Transaction business
func NewTransactionBusiness(TransactionRepository transaction.ITransactionRepository, TransactionService transaction_service.ITransactionService) ITransactionBusiness {
	return &TransactionBusiness{TransactionRepository, TransactionService}
}

// PostTransaction : post transaction
func (a *TransactionBusiness) PostTransaction(transaction models.Transaction) (int, *models.Transaction, error) {

	userPayer, err := a.TransactionRepository.GetUserById(transaction.Payer_id)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusInternalServerError, nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusBadRequest, nil, errors.New("payer not found")
	}

	if config.IsCNPJ(userPayer.CPFCNPJ) {
		return http.StatusBadRequest, nil, errors.New("payer with CNPJ can only recive payments")
	}

	userPayerWallet, err := a.TransactionRepository.GetUserWalletByUserId(userPayer.Id)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusInternalServerError, nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusBadRequest, nil, errors.New("there is no wallet registered for this payer_id")
	}

	if (userPayerWallet.Value < transaction.Value) || (userPayerWallet.Value == 0) {
		return http.StatusBadRequest, nil, errors.New("Insufficient balance to make transactions")
	}

	userReceive, err := a.TransactionRepository.GetUserById(transaction.Payer_receive_id)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusInternalServerError, nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusBadRequest, nil, errors.New("payer receive not found")
	}

	userPayerReceiveWallet, err := a.TransactionRepository.GetUserWalletByUserId(userReceive.Id)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusInternalServerError, nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusBadRequest, nil, errors.New("there is no wallet registered for this payer_receive_id")
	}

	response, err := a.TransactionService.CallApi("https://run.mocky.io/v3/d02168c6-d88d-4ff2-aac6-9e9eb3425e31")

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	if !response.Authorization {
		return http.StatusInternalServerError, nil, errors.New("unauthorized transaction")
	}

	err = a.TransactionRepository.SubtractValueFromWallet(userPayerWallet, transaction.Value)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	err = a.TransactionRepository.AddValueToWallet(userPayerReceiveWallet, transaction.Value)

	if err != nil {
		a.TransactionRepository.AddValueToWallet(userPayerWallet, transaction.Value)

		return http.StatusInternalServerError, nil, err
	}

	NewTransaction, err := a.TransactionRepository.Create(&models.Transaction{
		Value:            transaction.Value,
		Payer_id:         transaction.Payer_id,
		Payer_receive_id: transaction.Payer_receive_id,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	})

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusAccepted, NewTransaction, nil
}

// GetTransaction : get transaction from database by ID
func (a *TransactionBusiness) GetTransaction(id int64) (int, *[]models.Transaction, error) {
	transaction, err := a.TransactionRepository.FindById(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusBadRequest, nil, errors.New("transaction not found")
	}

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, transaction, nil
}

// GetTransactions : get all transactions from database
func (a *TransactionBusiness) GetTransactions() (int, *[]models.Transaction, error) {
	transaction, err := a.TransactionRepository.FindAll()

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusBadRequest, nil, errors.New("transactions not found")
	}

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, transaction, nil
}
