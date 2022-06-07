package user

import (
	"errors"
	"net/http"
	"time"

	"github.com/JoaoGabrielManochio/webapi-go/config"
	"github.com/JoaoGabrielManochio/webapi-go/models"
	"github.com/JoaoGabrielManochio/webapi-go/repository/user"
	"gorm.io/gorm"
)

// IUserBusiness : interface of user business
type IUserBusiness interface {
	PostUser(user models.User) (int, *models.User, error)
	// GetUsers(userId string) (int, *model.DocumentStatusInfo, error)
}

// UserBusiness : struct of user business
type UserBusiness struct {
	UserRepository user.IUserRepository
}

// NewUserBusiness : create a new user business
func NewUserBusiness(UserRepository user.IUserRepository) IUserBusiness {
	return &UserBusiness{UserRepository}
}

// PostUser : post user
func (a *UserBusiness) PostUser(user models.User) (int, *models.User, error) {

	_, err := a.UserRepository.GetUserByEmail(user.Email)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusInternalServerError, nil, err
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusBadRequest, nil, errors.New("email already used")
	}

	_, err = a.UserRepository.GetUserByDocument(user.CPFCNPJ)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusInternalServerError, nil, err
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusBadRequest, nil, errors.New("CPF/CNPJ already used")
	}

	if !config.IsCPF(user.CPFCNPJ) && !config.IsCNPJ(user.CPFCNPJ) {
		return http.StatusBadRequest, nil, errors.New("CPF/CNPJ is not valid")
	}

	newUser, err := a.UserRepository.Create(&models.User{
		Name:      user.Name,
		Email:     user.Email,
		CPFCNPJ:   user.CPFCNPJ,
		Password:  user.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusAccepted, newUser, nil
}

/*

// GetDocumentStatus : call bankly service and get document status with document type
func (a *UserBusiness) GetDocumentStatus(userId string) (int, *model.DocumentStatusInfo, error) {
	documents, err := a.UserRepository.Get(&model.UserDocument{
		UserId:   userId,
		IsActive: true,
	})

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	documentsStatus := getDocumentsStatus(documents)
	return http.StatusOK, documentsStatus, nil
}
*/
