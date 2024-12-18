package domain

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	ID          int             `json:"id" gorm:"primaryKey"`
	Name        string          `json:"name" gorm:"size:255;null"`
	Email       string          `json:"username" gorm:"size:255;null"`
	Password    string          `json:"password" gorm:"size:255;null"`
	PhoneNumber string          `json:"phone_number" gorm:"size:255;null"`
	CreatedAt   *time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   *time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (UserModel) TableName() string {
	return "users"
}
