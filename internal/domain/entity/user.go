package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role string

const (
  Admin     Role = "admin"
  Member    Role = "member"
)

type User struct {
	gorm.Model
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;"`
	Name  	  string         `json:"name,omitempty" gorm:"varchar(255);not null"`
	Username  string         `json:"username,omitempty" gorm:"not null"`
	Email     string         `json:"email,omitempty" gorm:"uniqueIndex;not null"`
	Password  string         `json:"password,omitempty" gorm:"not null"`
	Phone     string         `json:"phone,omitempty" gorm:"not null"`
	IsActive  bool           `json:"is_active" gorm:"not null;default:false"`
	Role      Role         	`json:"role,omitempty" gorm:"enum('admin', 'member');default:member"`
	Token     string         `json:"token"`
	Otp       string         `json:"otp"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}
