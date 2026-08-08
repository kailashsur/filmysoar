package models

import (
	"time"
)

type TrendingMovie struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	MovieID   int       `gorm:"column:movieId;unique;not null" json:"movieId"`
	Movie     Movie     `gorm:"foreignKey:MovieID;constraint:OnDelete:CASCADE" json:"movie,omitempty"`
	Order     int       `gorm:"default:0" json:"order"`
	CreatedAt time.Time `gorm:"column:createdAt;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt;autoUpdateTime" json:"updatedAt"`
}

func (TrendingMovie) TableName() string {
	return "trending_movies"
}
