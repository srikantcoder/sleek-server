package entities

import (
	"github.com/google/uuid"
)

type DealStatus struct {
	ID        uuid.UUID `json:"deal_id"`
	UserCount uint      `gorm:"not null" json:"user_count"`
}
