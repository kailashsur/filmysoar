package models

import (
	"time"
)

type Category struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"unique;not null" json:"name"`
	Slug        string    `gorm:"unique;not null" json:"slug"`
	Description *string   `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `gorm:"column:createdAt;autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updatedAt;autoUpdateTime" json:"updatedAt"`
	Movies      []Movie   `gorm:"foreignKey:CategoryID" json:"movies,omitempty"`
}

func (Category) TableName() string {
	return "categories"
}

type CategoryWithCount struct {
	Category
	MovieCount int `json:"_count"`
}
