package models

import "github.com/google/uuid"

type Session struct {
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID uint      `json:"user_id"`
}

func (s Session) ISDefault() bool {
	return s.UserID == 0 || len(s.ID) == 0 
}
