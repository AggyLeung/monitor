package model

import (
	"time"

	"github.com/google/uuid"
)

type Relation struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	SourceCIID uuid.UUID `gorm:"column:source_ci_id;type:uuid;index" json:"source_ci_id"`
	TargetCIID uuid.UUID `gorm:"column:target_ci_id;type:uuid;index" json:"target_ci_id"`
	Type       string    `gorm:"column:type;size:50" json:"type"`
	Properties []byte    `gorm:"type:jsonb" json:"properties"`
	CreatedAt  time.Time `json:"created_at"`
}

func (Relation) TableName() string {
	return "relation"
}
