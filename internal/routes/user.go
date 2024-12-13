package routes

import (
	"github.com/labstack/echo/v4"
	"mobile-auth/internal/handlers"
	"mobile-auth/internal/repositories"
)

func User(e *echo.Echo) {
	repo := repositories.NewUserRepository()
	handler := handlers.NewUserHandler(*repo)

	e.POST("/register", handler.Register)
	//e.POST("/auth", handler.Create, middleware.RequireRole("admin"))
	//e.POST("/activate", handler.Create, middleware.RequireRole("admin"))
	//e.POST("/confirm/:token", handler.Create, middleware.RequireRole("admin"))
	//e.POST("/restore", handler.Create, middleware.RequireRole("admin"))
}
