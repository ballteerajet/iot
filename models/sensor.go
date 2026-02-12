package models

import "time"

type SensorData struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`

	Temperature float64 `gorm:"not null" json:"temp"`
	Humidity    float64 `gorm:"not null" json:"humidity"`
}
