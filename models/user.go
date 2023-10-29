package models

import (
	"time"
)

type User struct {
	Id        uint   `gorm:"primaryKey" json:"id_user"`
	Name      string `json:"name_user" validate:"required"`
	Email     string `gorm:"unique_index;not null" json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAdt *time.Time
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
