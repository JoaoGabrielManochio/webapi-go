package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint           `gorm:"column:id;primary_key" json:"id"`
	Name      string         `gorm:"column:name;type:varchar(150);not null" json:"name" validate:"required"`
	Email     string         `gorm:"column:email;type:varchar(200);not null" json:"email" validate:"required"`
	Password  string         `gorm:"column:password;type:varchar(150);not null" json:"password" validate:"required"`
	CPFCNPJ   string         `gorm:"column:cpf_cnpj;type:varchar(150);not null" json:"cpf_cnpj" validate:"required"`
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
}
