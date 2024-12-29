package contacts

import (
	"gorm.io/gorm"
)

type Contacts struct {
	*gorm.Model
	Name         string `gorm:"size:255;not null"`
	Email        string `gorm:"size:255;unique;not null"`
	Role         string `gorm:"size:50;not null"` // e.g., Admin, KAM
	PhoneNo      string `gorm:"size:20"`
	RestaurantID uint   `gorm:"not null;index"`
}

func (u *Contacts) TableName() string {
	return "contacts"
}
