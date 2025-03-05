package models

import (
	"time"

	"github.com/google/uuid"
)

type Context string

const (
	Login          Context = "Login"
	Logout         Context = "Logout"
	ForgotPassword Context = "Forgot Password"
	UpdatePassword Context = "Update Password"
	Order          Context = "Create Order"
	Payment        Context = "Captured Payment"
)

type UsersActivityHistory struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	UserID      uuid.UUID `json:"user_id" gorm:"not null"`
	Context     Context   `json:"context" gorm:"varchar(255);not null"`
	Ip          string    `json:"ip" gorm:"varchar(255);not null"`
	AgentClient string    `json:"agent_client" gorm:"varchar(255);not null"`
	Device      string    `json:"device" gorm:"varchar(255);not null"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	User        User      `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
}

func (UsersActivityHistory) TableName() string {
	return "users_activity_histories"
}
