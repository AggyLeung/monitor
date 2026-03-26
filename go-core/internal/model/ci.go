package model

import (
	"net"
	"time"

	"github.com/google/uuid"
)

type CI struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	CITypeID  int        `gorm:"column:ci_type_id;index" json:"ci_type_id"`
	Name      string     `gorm:"size:255;not null" json:"name"`
	Status    string     `gorm:"size:50;not null;default:active" json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	CreatedBy *uuid.UUID `gorm:"type:uuid" json:"created_by,omitempty"`
	UpdatedBy *uuid.UUID `gorm:"type:uuid" json:"updated_by,omitempty"`
}

// ServerCI extends CI and uses embedded fields for joined table style persistence.
type ServerCI struct {
	CI     `gorm:"embedded"`
	CPU    int    `json:"cpu"`
	Memory int    `json:"memory"`
	IP     net.IP `gorm:"type:inet" json:"ip"`
	OS     string `gorm:"size:64" json:"os"`
}

func (CI) TableName() string {
	return "ci"
}

type CIAttributeValue struct {
	CIID        uuid.UUID `gorm:"column:ci_id;type:uuid;primaryKey" json:"ci_id"`
	AttributeID int       `gorm:"column:attribute_id;primaryKey" json:"attribute_id"`
	Value       []byte    `gorm:"type:jsonb" json:"value"`
}

func (CIAttributeValue) TableName() string {
	return "ci_attribute_value"
}
