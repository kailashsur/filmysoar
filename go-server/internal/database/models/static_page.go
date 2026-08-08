package models

import (
	"time"
)

type StaticPage struct {
	ID              int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title           string    `gorm:"not null" json:"title"`
	Slug            string    `gorm:"unique;not null" json:"slug"`
	Content         string    `gorm:"type:text;not null" json:"content"`
	MetaTitle       *string   `gorm:"type:text" json:"metaTitle"`
	MetaDescription *string   `gorm:"type:text" json:"metaDescription"`
	MetaKeywords    *string   `gorm:"type:text" json:"metaKeywords"`
	IsPublished     bool      `gorm:"default:true" json:"isPublished"`
	CreatedAt       time.Time `gorm:"column:createdAt;autoCreateTime" json:"createdAt"`
	UpdatedAt       time.Time `gorm:"column:updatedAt;autoUpdateTime" json:"updatedAt"`
}

func (StaticPage) TableName() string {
	return "static_pages"
}
