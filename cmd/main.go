package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"mobile-auth/internal/db"
	"mobile-auth/internal/routes"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Env load error: %v\n", err)
	}

	if err := os.MkdirAll(os.Getenv("LOG_DIR"), os.ModePerm); err != nil {
		log.Fatalf("Не удалось создать директорию для логов: %v", err)
	}

	filePath := filepath.Join(os.Getenv("LOG_DIR"), os.Getenv("LOG_FILE"))
	logFileHandle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Не удалось открыть файл для логов: %v", err)
	}
	defer logFileHandle.Close()

	log.SetOutput(logFileHandle)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Mobile Auth API")
	})

	db.ConnectDatabase()

	routes.User(e)

	e.Logger.Fatal(e.Start(os.Getenv("ECHO_HOST")))

}
