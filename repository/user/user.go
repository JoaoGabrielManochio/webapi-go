package user

import (
	"github.com/JoaoGabrielManochio/webapi-go/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(user *models.User, id int64) error
	FindById(id int64) (*[]models.User, error)
	FindAll() (*[]models.User, error)
	GetUserByEmail(email string) (*[]models.User, error)
	GetUserByEmailAndId(email string, id uint) (*[]models.User, error)
	GetUserByDocument(cpf_cnpj string) (*[]models.User, error)
	GetUserByDocumentAndId(cpf_cnpj string, id uint) (*[]models.User, error)
}

// User : struct of user repository
type UserRepository struct{ db *gorm.DB }

// NewUser : create a new user repository
func NewUser(db *gorm.DB) IUserRepository {
	return &UserRepository{db}
}

// -> verificar tirar a senha no retorno da api
// FindById : get user by ID
func (b *UserRepository) FindById(id int64) (*[]models.User, error) {
	user := &[]models.User{}
	find := b.db

	if id != 0 {
		find = find.Where("id = ?", id)
	}

	err := find.
		First(user).
		Error

	return user, err
}

// FindAll : get all users
func (b *UserRepository) FindAll() (*[]models.User, error) {
	user := &[]models.User{}
	find := b.db

	err := find.
		Find(user).
		Error

	return user, err
}

// Create : create user
func (b *UserRepository) Create(user *models.User) (*models.User, error) {

	create := b.db

	err := create.
		Create(user).
		Error

	return user, err
}

// Update : update user
func (b *UserRepository) Update(user *models.User) (*models.User, error) {

	create := b.db

	err := create.
		Save(user).
		Error

	return user, err
}

// Delete : delete user by ID
func (b *UserRepository) Delete(user *models.User, id int64) error {

	create := b.db

	err := create.
		Delete(user, id).
		Error

	return err
}

// GetUserByEmail : get user by email
func (b *UserRepository) GetUserByEmail(email string) (*[]models.User, error) {

	user := &[]models.User{}
	find := b.db

	if email != "" {
		find = find.Where("email = ?", email)
	}

	err := find.First(user)

	return user, err.Error
}

// GetUserByDocument : get user by document
func (b *UserRepository) GetUserByDocument(cpf_cnpj string) (*[]models.User, error) {

	user := &[]models.User{}
	find := b.db

	if cpf_cnpj != "" {
		find = find.Where("cpf_cnpj = ?", cpf_cnpj)
	}

	err := find.First(user)

	return user, err.Error
}

// GetUserByEmailAndId : get user by email where ID is different from id
func (b *UserRepository) GetUserByEmailAndId(email string, id uint) (*[]models.User, error) {

	user := &[]models.User{}
	find := b.db

	if email != "" {
		find = find.Where("email = ? AND id <> ?", email, id)
	}

	err := find.First(user)

	return user, err.Error
}

// GetUserByDocument : get user by document where ID is different from ID
func (b *UserRepository) GetUserByDocumentAndId(cpf_cnpj string, id uint) (*[]models.User, error) {

	user := &[]models.User{}
	find := b.db

	if cpf_cnpj != "" {
		find = find.Where("cpf_cnpj = ? AND id <> ?", cpf_cnpj, id)
	}

	err := find.First(user)

	return user, err.Error
}
