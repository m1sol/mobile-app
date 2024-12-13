package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"mobile-auth/internal/entities"
	"mobile-auth/internal/repositories"
	"mobile-auth/internal/responses"
	"mobile-auth/internal/utils"
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
	user.ConfirmationToken = token

	if err := h.Repo.Register(&user); err != nil {
		log.Printf("Register Error: %s", err.Error())
		return responses.InternalServerErrorResponse(c, err)
	}

	return responses.NoContentResponse(c)
}

//func (h *UserHandler) Auth(c echo.Context) error {
//}
//
//func (h *UserHandler) Activate(c echo.Context) error {
//}
//
//func (h *UserHandler) Confirm(c echo.Context) error {
//}
//
//func (h *UserHandler) Restore(c echo.Context) error {
//}
