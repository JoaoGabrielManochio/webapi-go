package user

import (
	"github.com/JoaoGabrielManochio/webapi-go/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	FindById(id int64) (*[]models.User, error)
	FindAll() (*[]models.User, error)
}

// User : struct of bank repository
type UserRepository struct{ db *gorm.DB }

// NewUser : create a new bank repository
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

func (b *UserRepository) Create() (*[]models.User, error) {
	user := &[]models.User{}
	create := b.db

	err := create.
		Create(user).
		Error

	return user, err
}
