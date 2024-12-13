package entities

import "time"

type User struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	Login             string    `gorm:"size:50;not null;uniqueIndex" json:"login"`
	Name              string    `gorm:"size:100;not null" json:"name"`
	Email             string    `gorm:"size:255;not null;uniqueIndex" json:"email"`
	Password          string    `gorm:"size:255;not null" json:"password"`
	EmailConfirmed    bool      `gorm:"default:false" json:"email_confirmed"`
	ConfirmationToken string    `gorm:"size:255;uniqueIndex" json:"confirmationToken"`
	TokenExpiry       time.Time `gorm:"index" json:"tokenExpiry"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
