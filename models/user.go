package models

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"iot/config"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
	Role     string `gorm:"default:user" json:"role"`
	APIKey   string `gorm:"unique;index" json:"api_key"`
}

func GenerateAPIKey() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func InitAdmin() {
	var count int64
	config.DB.Model(&User{}).Where("role = ?", "admin").Count(&count)

	if count == 0 {
		admin := User{
			Username: "admin",
			Password: "admin123",
			Role:     "admin",
			APIKey:   GenerateAPIKey(),
		}
		config.DB.Create(&admin)
	}
}
