package interactions

import (
	"github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
	"time"
)

type Interaction struct {
	*gorm.Model
	RestaurantID        uint           `gorm:"not null;index" json:"restaurant_id"`
	ContactID           uint           `gorm:"not null;index" json:"contact_id"`
	LeadID              uint           `gorm:"not null;index" json:"lead_id"`
	InteractionDate     time.Time      `gorm:"not null" json:"interaction_date"`
	InteractionType     string         `gorm:"size:50;not null" json:"interaction_type"` // e.g., Call, Meeting
	Details             postgres.Jsonb `gorm:"type:jsonb" json:"details"`
	LastInteractionDate time.Time      `gorm:"not null;index" json:"last_interaction_date"`
	NextInteractionDate time.Time      `gorm:"not null;index" json:"next_interaction_date"`
}

func (i *Interaction) TableName() string {
	return "interactions"
}
