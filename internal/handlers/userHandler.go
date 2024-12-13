package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"log"
	"mobile-auth/internal/entities"
	"mobile-auth/internal/repositories"
	"mobile-auth/internal/responses"
	"mobile-auth/internal/utils"
	"time"
)

type UserHandler struct {
	Repo repositories.UserRepository
}

func NewUserHandler(repo repositories.UserRepository) *UserHandler {
	return &UserHandler{Repo: repo}
}

func (h *UserHandler) Register(c echo.Context) error {
	var user entities.User
	if err := c.Bind(&user); err != nil {
		log.Printf("Register Error: %s", err.Error())
		return responses.InternalServerErrorResponse(c, err)
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		log.Printf("Register Error: %s", err.Error())
		return responses.InternalServerErrorResponse(c, err)
	}

	passwordHash, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Printf("Register Error: %s", err.Error())
		return responses.InternalServerErrorResponse(c, err)
	}
	user.Password = passwordHash

	token, err := utils.GenerateToken()
	if err != nil {
		log.Printf("Register Error: %s", err.Error())
		return responses.InternalServerErrorResponse(c, err)
	}
	user.ConfirmToken = token
	user.TokenExpiry = time.Now().Add(24 * time.Hour)

	if err := h.Repo.RegisterWithEmail(&user); err != nil {
		log.Printf("Register Error: %s", err.Error())
		return responses.InternalServerErrorResponse(c, err)
	}
	return responses.NoContentResponse(c)
}

func (h *UserHandler) Confirm(c echo.Context) error {

	token := c.Param("token")
	if err := h.Repo.Confirm(token); err != nil {
		log.Printf("Confirm Error: %s", err.Error())
		return responses.InternalServerErrorResponse(c, err)
	}
	return responses.NoContentResponse(c)
}

func (h *UserHandler) ResendConfirm(c echo.Context) error {
	var input struct {
		Email string `json:"email"`
	}
	if err := c.Bind(&input); err != nil {
		log.Printf("ResendConfirm Error: %s", err.Error())
		return responses.InternalServerErrorResponse(c, err)
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		log.Printf("ResendConfirm Error: %s", err.Error())
		return responses.InternalServerErrorResponse(c, err)
	}

	if err := h.Repo.CheckUserByEmail(input.Email); err != nil {
		log.Printf("ResendConfirm Error: %s", err.Error())
		return responses.InternalServerErrorResponse(c, err)
	}

	token, err := utils.GenerateToken()
	if err != nil {
		log.Printf("Register Error: %s", err.Error())
		return responses.InternalServerErrorResponse(c, err)
	}

	if err := h.Repo.ResendToken(input.Email, token); err != nil {
		log.Printf("Register Error: %s", err.Error())
		return responses.InternalServerErrorResponse(c, err)
	}

	return responses.NoContentResponse(c)
}
