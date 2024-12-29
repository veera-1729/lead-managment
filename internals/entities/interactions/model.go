package interactions

import (
	"github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
	"time"
)

type Interaction struct {
	*gorm.Model
	RestaurantID    uint           `gorm:"not null;index"`
	ContactID       uint           `gorm:"not null;index"`
	LeadID          uint           `gorm:"not null;index;F"`
	InteractionDate time.Time      `gorm:"not null"`
	InteractionType string         `gorm:"size:50;not null"` // e.g., Call, Meeting
	Details         postgres.Jsonb `gorm:"type:text"`
}
