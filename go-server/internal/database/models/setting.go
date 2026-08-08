package models

import (
	"time"
)

type Setting struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Key         string    `gorm:"unique;not null" json:"key"`
	Value       string    `gorm:"type:text;not null" json:"value"`
	Description *string   `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `gorm:"column:createdAt;autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updatedAt;autoUpdateTime" json:"updatedAt"`
}

func (Setting) TableName() string {
	return "settings"
}
