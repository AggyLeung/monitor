package model

import (
	"time"

	"github.com/google/uuid"
)

type SyncTask struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	TaskType    string     `gorm:"column:task_type;size:50" json:"task_type"`
	Status      string     `gorm:"size:20" json:"status"`
	Params      []byte     `gorm:"type:jsonb" json:"params"`
	Result      []byte     `gorm:"type:jsonb" json:"result"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `gorm:"column:completed_at" json:"completed_at,omitempty"`
}

func (SyncTask) TableName() string {
	return "sync_task"
}
