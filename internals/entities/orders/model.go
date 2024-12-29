package orders

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	*gorm.Model
	RestaurantID uint      `gorm:"not null;index" json:"restaurant_id"`
	OrderDate    time.Time `gorm:"not null" json:"order_date"`
	Amount       float64   `gorm:"not null" json:"amount"`
	Status       string    `gorm:"size:50;not null" json:"status"` // e.g., Completed, Pending
}
