package main

import (
	"fmt"
	"time"

	// "github.com/Hdeee1/be-go-restaurant-management/internal/delivery/handler"
	"github.com/Hdeee1/be-go-restaurant-management/internal/repository/mysql"
	"github.com/Hdeee1/be-go-restaurant-management/internal/services"
	"github.com/Hdeee1/be-go-restaurant-management/pkg/database"
	// "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("failed to load .env file")
	}

	db, err := database.ConnectDB()
	if err != nil {
		fmt.Errorf("failed to connect to database, error: %w", err.Error())
	}

	repo := mysql.NewMySQLRepository(db)
	_ = services.NewUserServices(repo, 5 * time.Minute)
	// handler := handler.NewUserHandler()
}