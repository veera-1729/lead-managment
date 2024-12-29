package orders

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	*gorm.Model
	RestaurantID uint      `gorm:"not null;index"`
	OrderDate    time.Time `gorm:"not null"`
	Amount       float64   `gorm:"not null"`
	Status       string    `gorm:"size:50;not null"` // e.g., Completed, Pending
}
