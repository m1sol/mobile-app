package repositories

import (
	"mobile-auth/internal/db"
	"mobile-auth/internal/entities"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) Register(user *entities.User) error {
	//bcrypt.GenerateFromPassword()
	if err := db.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
