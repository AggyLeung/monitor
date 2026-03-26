package model

import "time"

type CIType struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:100;uniqueIndex;not null" json:"name"`
	Label     string    `gorm:"size:100" json:"label"`
	Icon      string    `gorm:"size:50" json:"icon"`
	CreatedAt time.Time `json:"created_at"`
}

func (CIType) TableName() string {
	return "ci_type"
}

type CITypeAttribute struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	CITypeID    int    `gorm:"not null;index" json:"ci_type_id"`
	Name        string `gorm:"size:100;not null" json:"name"`
	Label       string `gorm:"size:100" json:"label"`
	DataType    string `gorm:"size:50;not null" json:"data_type"`
	IsRequired  bool   `gorm:"default:false" json:"is_required"`
	EnumOptions []byte `gorm:"type:jsonb" json:"enum_options"`
	SortOrder   int    `json:"sort_order"`
}

func (CITypeAttribute) TableName() string {
	return "ci_type_attribute"
}
