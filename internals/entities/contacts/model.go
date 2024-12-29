package contacts

import (
	"gorm.io/gorm"
)

type Contacts struct {
	*gorm.Model
	Name         string `gorm:"size:255;not null" json:"name"`
	Email        string `gorm:"size:255;unique;not null" json:"email"`
	Role         string `gorm:"size:50;not null" json:"role"` // e.g., Admin, KAM
	PhoneNo      string `gorm:"size:20;unique;not null" json:"phone_no"`
	RestaurantID uint   `gorm:"not null;index" json:"restaurant_id"`
}

func (u *Contacts) TableName() string {
	return "contacts"
}
