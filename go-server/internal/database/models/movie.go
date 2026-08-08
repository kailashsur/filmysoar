package models

import (
	"time"
)

type Movie struct {
	ID          int            `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Slug        string         `gorm:"unique;not null" json:"slug"`
	Description *string        `gorm:"type:text" json:"description"`
	Thumbnail   *string        `gorm:"type:text" json:"thumbnail"`
	Genre       *string        `json:"genre"`
	Languages   *string        `json:"languages"`
	Duration    *string        `json:"duration"`
	ReleaseYear *int           `gorm:"column:releaseYear" json:"releaseYear"`
	Cast        *string        `json:"cast"`
	Sizes       *string        `json:"sizes"`
	DownloadURL *string        `gorm:"column:downloadUrl;type:text" json:"downloadUrl"`
	Screenshot  *string        `gorm:"type:text" json:"screenshot"`
	Keywords    *string        `json:"keywords"`
	CategoryID  *int           `gorm:"column:categoryId" json:"categoryId"`
	Category    *Category      `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	CreatedAt   time.Time      `gorm:"column:createdAt;autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"column:updatedAt;autoUpdateTime" json:"updatedAt"`
	Trending    *TrendingMovie `gorm:"foreignKey:MovieID" json:"TrendingMovie,omitempty"`
}

func (Movie) TableName() string {
	return "movies"
}

type MovieWithTrending struct {
	Movie
	IsTrending bool `json:"isTrending"`
}
