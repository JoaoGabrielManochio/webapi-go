package user

import (
	"github.com/JoaoGabrielManochio/webapi-go/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user *models.User) (*models.User, error)
	FindById(id int64) (*[]models.User, error)
	FindAll() (*[]models.User, error)
	GetUserByEmail(email string) (*[]models.User, error)
	GetUserByDocument(cpf_cnpj string) (*[]models.User, error)
}

// User : struct of user repository
type UserRepository struct{ db *gorm.DB }

// NewUser : create a new user repository
func NewUser(db *gorm.DB) IUserRepository {
	return &UserRepository{db}
}

func (b *UserRepository) FindById(id int64) (*[]models.User, error) {
	user := &[]models.User{}
	find := b.db

	if id != 0 {
		find = find.Where("id = ?", id)
	}

	err := find.
		Find(user).
		Error

	return user, err
}

func (b *UserRepository) FindAll() (*[]models.User, error) {
	user := &[]models.User{}
	find := b.db

	err := find.
		Find(user).
		Error

	return user, err
}

func (b *UserRepository) Create(user *models.User) (*models.User, error) {

	create := b.db

	err := create.
		Create(user).
		Error

	return user, err
}

func (b *UserRepository) GetUserByEmail(email string) (*[]models.User, error) {

	user := &[]models.User{}
	find := b.db

	if email != "" {
		find = find.Where("email = ?", email)
	}

	err := find.First(user)

	return user, err.Error
}

func (b *UserRepository) GetUserByDocument(cpf_cnpj string) (*[]models.User, error) {

	user := &[]models.User{}
	find := b.db

	if cpf_cnpj != "" {
		find = find.Where("cpf_cnpj = ?", cpf_cnpj)
	}

	err := find.First(user)

	return user, err.Error
}
