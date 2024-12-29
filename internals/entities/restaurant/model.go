package restaurant

import (
	"github.com/veera-1729/lead-managment/internals/entities/contacts"
	"github.com/veera-1729/lead-managment/internals/entities/interactions"
	"github.com/veera-1729/lead-managment/internals/entities/orders"
	"gorm.io/gorm"
)

type Restaurant struct {
	*gorm.Model
	Name         string                     `gorm:"size:255;not null" json:"name"`
	Address      string                     `gorm:"size:500" json:"address"`
	Status       string                     `gorm:"size:50;not null" json:"status"` // e.g., New, In Progress, Converted
	Contacts     []contacts.Contacts        `gorm:"foreignKey:RestaurantID" json:"contacts"`
	Orders       []orders.Order             `gorm:"foreignKey:RestaurantID" json:"orders"`
	Interactions []interactions.Interaction `gorm:"foreignKey:RestaurantID" json:"interactions"`
	LeadID       string                     `gorm:"size:50" json:"leadID"`
}

func (r *Restaurant) TableName() string {
	return "restaurants"
}
