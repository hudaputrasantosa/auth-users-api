package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;"`
	Username  string         `json:"username" validate:"required" gorm:"not null"`
	Email     string         `json:"email" validate:"required,email" gorm:"not null"`
	Password  string         `json:"password" validate:"required" gorm:"not null, colum:password"`
	Phone     uint           `json:"phone" validate:"required,number,min=12" gorm:"required,not null"`
	IsActive  bool           `json:"is_active" gorm:"not null, default:false"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Users struct {
	Users []User `json:"users"`
}

// func (user *User) beforeCreate(tx *gorm.DB) (err error) {
// 	user.ID = uuid.New().String()
// 	return
// }
