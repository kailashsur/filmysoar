package models

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:text;not null" json:"name"`
	Email     string    `gorm:"type:text;not null" json:"email"`
	Password  string    `gorm:"type:text;not null" json:"password"`
	Role      string    `gorm:"type:text;default:user" json:"role"`
	Status    string    `gorm:"type:text;default:active" json:"status"`
	CreatedAt time.Time `gorm:"column:createdAt;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt;autoUpdateTime" json:"updatedAt"`
}

func (User) TableName() string {
	return "users"
}
