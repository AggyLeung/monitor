package model

import (
	"time"

	"github.com/google/uuid"
)

type GraphSyncFailed struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	EntityType   string     `gorm:"column:entity_type;size:50;not null" json:"entity_type"`
	EntityID     *uuid.UUID `gorm:"column:entity_id;type:uuid" json:"entity_id,omitempty"`
	Payload      []byte     `gorm:"type:jsonb;not null" json:"payload"`
	ErrorMessage string     `gorm:"column:error_message;type:text" json:"error_message"`
	Status       string     `gorm:"size:20;not null;default:pending" json:"status"`
	RetryCount   int        `gorm:"column:retry_count;not null;default:0" json:"retry_count"`
	NextRetryAt  *time.Time `gorm:"column:next_retry_at" json:"next_retry_at,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

func (GraphSyncFailed) TableName() string {
	return "graph_sync_failed"
}
