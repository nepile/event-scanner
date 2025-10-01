package entity

import "github.com/google/uuid"

type Presence struct {
	Base
	RegistrantID *uuid.UUID `gorm:"type:char(36)" json:"registrant_id"`
	Registrant   Registrant `gorm:"foreignKey:RegistrantID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
