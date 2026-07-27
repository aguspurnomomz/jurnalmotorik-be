package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Username       string    `gorm:"type:varchar(50);unique;not null" json:"username"`
	Fullname       string    `gorm:"type:varchar(100);not null" json:"fullname"`
	Email          string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	PasswordHash   string    `gorm:"type:varchar(255);not null" json:"-"` // "-" artinya tidak dikembalikan di respon JSON
	WhatsappNumber string    `gorm:"type:varchar(20)" json:"whatsapp_number"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	Children []Child `gorm:"foreignKey:UserID" json:"children,omitempty"`
}