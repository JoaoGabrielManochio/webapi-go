package models

import (
	"time"

	"gorm.io/gorm"
)

type Wallet struct {
	Id        uint           `gorm:"column:id;primary_key" json:"id"`
	Name      string         `gorm:"column:name;type:varchar(150);not null" json:"name" validate:"required"`
	UserId    int64          `gorm:"column:user_id;type:int(100);not null" json:"user_id" validate:"required"`
	Value     float64        `gorm:"column:value;type:float(10,2);not null" json:"value" validate:"required"`
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
}
