package models

import (
	"time"

	"github.com/google/uuid"
)

type Child struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	BirthDate string    `gorm:"type:date;not null" json:"birth_date"`
	Gender    string    `gorm:"type:varchar(1);not null" json:"gender"`
	HeightCm  float64   `gorm:"type:numeric(5,2)" json:"height_cm"`
	WeightKg  float64   `gorm:"type:numeric(5,2)" json:"weight_kg"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}