package entity

import "github.com/google/uuid"

type Registrant struct {
	Base
	UserID  *uuid.UUID `gorm:"type:char(36)" json:"user_id"`
	EventID *uuid.UUID `gorm:"type:char(36)" json:"event_id"`
	User    User       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Event   Event      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
