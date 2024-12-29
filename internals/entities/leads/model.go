package leads

import (
	"gorm.io/gorm"
)

type Lead struct {
	*gorm.Model
	RestaurantID  uint   `gorm:"not null;index" json:"restaurant_id"`
	Name          string `gorm:"size:255;not null" json:"name"`
	Role          string `gorm:"size:50" json:"role"` // e.g., Manager, Chef
	Phone         string `gorm:"size:20" json:"phone"`
	Email         string `gorm:"size:255" json:"email"`
	Status        string `gorm:"size:50" json:"status"`           // e.g., Active, Contacted, Qualified, Lost
	CallFrequency uint   `gorm:"default:0" json:"call_frequency"` // e.g., 0, 1, 2, 3 in days
}

func (l *Lead) TableName() string {
	return "leads"
}
