package leads

import (
	"gorm.io/gorm"
)

type Lead struct {
	*gorm.Model
	RestaurantID  uint   `gorm:"not null;index"`
	Name          string `gorm:"size:255;not null"`
	Role          string `gorm:"size:50"` // e.g., Manager, Chef
	Phone         string `gorm:"size:20"`
	Email         string `gorm:"size:255"`
	Status        string `gorm:"size:50"`   // e.g., Active, Contacted, Qualified, Lost
	CallFrequency uint   `gorm:"default:0"` // e.g., 0, 1, 2, 3 in days
}
