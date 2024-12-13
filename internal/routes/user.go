package routes

import (
	"github.com/labstack/echo/v4"
	"mobile-auth/internal/handlers"
	"mobile-auth/internal/repositories"
	"mobile-auth/internal/utils"
	"os"
)

func User(e *echo.Echo) {
	emailSender := utils.NewEmailSender(
		os.Getenv("EMAIL_SMTP_SERVER"),
		os.Getenv("EMAIL_PORT"),
		os.Getenv("EMAIL_USERNAME"),
		os.Getenv("EMAIL_PASSWORD"),
	)
	//
	repo := repositories.NewUserRepository(emailSender)
	handler := handlers.NewUserHandler(*repo)

	e.POST("/register", handler.Register)
	e.GET("/confirm/:token", handler.Confirm)
	e.POST("/resend-token", handler.ResendConfirm)
}
