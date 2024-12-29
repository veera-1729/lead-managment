package restaurant

import (
	"github.com/veera-1729/lead-managment/internals/entities/contacts"
	"github.com/veera-1729/lead-managment/internals/entities/interactions"
	"github.com/veera-1729/lead-managment/internals/entities/orders"
	"gorm.io/gorm"
)

type Restaurant struct {
	*gorm.Model
	Name         string                     `gorm:"size:255;not null"`
	Address      string                     `gorm:"size:500"`
	Status       string                     `gorm:"size:50;not null"` // e.g., New, In Progress, Converted
	Contacts     []contacts.Contacts        `gorm:"foreignKey:RestaurantID"`
	Orders       []orders.Order             `gorm:"foreignKey:RestaurantID"`
	Interactions []interactions.Interaction `gorm:"foreignKey:RestaurantID"`
	LeadID       string                     `gorm:"size:50"`
}
