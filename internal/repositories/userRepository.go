package repositories

import (
	"errors"
	"fmt"
	"log"
	"mobile-auth/internal/db"
	"mobile-auth/internal/entities"
	"mobile-auth/internal/utils"
	"os"
	"time"
)

type UserRepository struct {
	EmailSender *utils.EmailSender
}

func NewUserRepository(emailSender *utils.EmailSender) *UserRepository {
	return &UserRepository{
		EmailSender: emailSender,
	}
}

func (u *UserRepository) RegisterWithEmail(user *entities.User) error {
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("RegisterWithEmail Panic: %v", r)
		}
	}()

	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	subject, body := generateEmailBody(user.Name, user.ConfirmToken)

	if err := u.EmailSender.SendEmail(user.Email, subject, body); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) Confirm(token string) error {
	var user entities.User
	if err := db.DB.Where(entities.User{ConfirmToken: token}).First(&user).Error; err != nil {
		return err
	}
	if user.TokenExpiry.Before(time.Now()) {
		return errors.New("token expired")
	}

	result := db.DB.Model(&user).
		Where(entities.User{ID: user.ID}).
		Select("EmailConfirmed", "ConfirmToken").
		Updates(map[string]interface{}{
			"EmailConfirmed": true,
			"ConfirmToken":   "",
		})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserRepository) CheckUserByEmail(email string) error {
	var user entities.User
	if err := db.DB.Where(entities.User{Email: email}).First(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) ResendToken(email, token string) error {
	var user entities.User

	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("RegisterWithEmail Panic: %v", r)
		}
	}()

	if err := tx.Where(entities.User{Email: email}).First(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	result := tx.Model(&user).
		Where(entities.User{ID: user.ID}).
		Select("ConfirmToken").
		Updates(map[string]interface{}{
			"ConfirmToken": token,
		})

	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	subject, body := generateEmailBody(user.Name, user.ConfirmToken)

	if err := u.EmailSender.SendEmail(user.Email, subject, body); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func generateEmailBody(name, token string) (string, string) {
	body := fmt.Sprintf(
		"Здравствуйте, %s!\n\nДля подтверждения вашего email, перейдите по ссылке:\nhttp://%s/confirm/%s",
		name,
		os.Getenv("ECHO_HOST"),
		token,
	)
	return "Подтверждение регистрации", body
}
