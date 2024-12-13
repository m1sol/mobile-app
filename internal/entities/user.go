package entities

import "time"

type User struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Login          string    `gorm:"size:50;not null;uniqueIndex" json:"login"`
	Name           string    `gorm:"size:100;not null" json:"name" validate:"required"`
	Email          string    `gorm:"size:255;not null;uniqueIndex" json:"email" validate:"required,email"`
	Password       string    `gorm:"size:255;not null" json:"password" validate:"required"`
	EmailConfirmed bool      `gorm:"default:false" json:"emailСonfirmed"`
	ConfirmToken   string    `gorm:"size:255;uniqueIndex" json:"-"`
	TokenExpiry    time.Time `gorm:"index" json:"-"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
