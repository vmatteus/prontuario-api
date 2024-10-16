package domain

import "time"

type UserModel struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	CreationDate *time.Time
	UpdateDate   *time.Time
}
