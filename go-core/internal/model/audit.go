package model

import (
	"time"

	"github.com/google/uuid"
)

type AuditLog struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID    *uuid.UUID `gorm:"column:user_id;type:uuid;index" json:"user_id,omitempty"`
	Action    string     `gorm:"size:50" json:"action"`
	CIID      *uuid.UUID `gorm:"column:ci_id;type:uuid;index" json:"ci_id,omitempty"`
	OldValue  []byte     `gorm:"column:old_value;type:jsonb" json:"old_value"`
	NewValue  []byte     `gorm:"column:new_value;type:jsonb" json:"new_value"`
	CreatedAt time.Time  `json:"created_at"`
}

func (AuditLog) TableName() string {
	return "audit_log"
}
