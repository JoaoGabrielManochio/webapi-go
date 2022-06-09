package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	Id               uint           `gorm:"column:id;primary_key" json:"id"`
	Payer_id         int64          `gorm:"column:payer_id;type:int(100);not null" json:"payer_id" validate:"required"`
	Payer_receive_id int64          `gorm:"column:payer_receive_id;type:int(100);not null" json:"payer_receive_id" validate:"required"`
	Value            float64        `gorm:"column:value;type:float(10,2);not null" json:"value" validate:"required"`
	CreatedAt        time.Time      `gorm:"column:created_at;type:datetime;not null" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;type:datetime;not null" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
}
