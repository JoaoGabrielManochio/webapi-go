package user

import (
	"crypto/md5"
	"encoding/hex"
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
	UpdateUser(user models.User) (int, *models.User, error)
	GetUser(id int64) (int, *[]models.User, error)
	GetUsers() (int, *[]models.User, error)
	DeleteUser(id int64) (int, error)
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

	password := createHash(user.Password)

	newUser, err := a.UserRepository.Create(&models.User{
		Name:      user.Name,
		Email:     user.Email,
		CPFCNPJ:   user.CPFCNPJ,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusAccepted, newUser, nil
}

// GetUser : get user from database by ID
func (a *UserBusiness) GetUser(id int64) (int, *[]models.User, error) {
	user, err := a.UserRepository.FindById(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusBadRequest, nil, errors.New("user not found")
	}

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, user, nil
}

// GetUsers : get all users from database
func (a *UserBusiness) GetUsers() (int, *[]models.User, error) {
	user, err := a.UserRepository.FindAll()

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusBadRequest, nil, errors.New("users not found")
	}

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, user, nil
}

// UpdateUser : update user infos
func (a *UserBusiness) UpdateUser(user models.User) (int, *models.User, error) {

	_, err := a.UserRepository.GetUserByEmailAndId(user.Email, user.Id)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusInternalServerError, nil, err
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusBadRequest, nil, errors.New("email already used")
	}

	_, err = a.UserRepository.GetUserByDocumentAndId(user.CPFCNPJ, user.Id)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusInternalServerError, nil, err
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusBadRequest, nil, errors.New("CPF/CNPJ already used")
	}

	if !config.IsCPF(user.CPFCNPJ) && !config.IsCNPJ(user.CPFCNPJ) {
		return http.StatusBadRequest, nil, errors.New("CPF/CNPJ is not valid")
	}

	newUser, err := a.UserRepository.Update(&models.User{
		Id:        user.Id,
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

// DeleteUser : delete user from database by ID
func (a *UserBusiness) DeleteUser(id int64) (int, error) {

	_, err := a.UserRepository.FindById(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusBadRequest, errors.New("user not found")
	}

	err = a.UserRepository.Delete(&models.User{}, id)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func createHash(key string) string {
	hasher := md5.New()

	hasher.Write([]byte(key))

	return hex.EncodeToString(hasher.Sum(nil))
}
