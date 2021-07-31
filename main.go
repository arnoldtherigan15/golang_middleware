package main

import (
	"fmt"
	"log"
	"os"

	_authService "section8/auth/service"
	_carHandler "section8/car/handler"
	_carRepo "section8/car/repository"
	_carService "section8/car/service"
	_garageHandler "section8/garage/handler"
	_garageRepo "section8/garage/repository"
	_garageService "section8/garage/service"
	_jwt "section8/jwt/service"
	_orderHandler "section8/order/handler"
	_orderRepo "section8/order/repository"
	_orderService "section8/order/service"
	_userHandler "section8/user/handler"
	_userRepo "section8/user/repository"
	_userService "section8/user/service"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file not found\n", err)
	}
}

func main() {
	DB_HOST := os.Getenv("DB_HOST")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_USER := os.Getenv("DB_USER")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Connection Database Error\n", err.Error())
	}

	authService := _authService.NewService()

	carRepo := _carRepo.NewRepository(db)
	carService := _carService.NewService(carRepo)

	garageRepo := _garageRepo.NewRepository(db)
	garageService := _garageService.NewService(garageRepo)

	userRepo := _userRepo.NewRepository(db)
	userService := _userService.NewService(userRepo, authService)

	orderRepo := _orderRepo.NewRepository(db)
	orderService := _orderService.NewService(orderRepo, carRepo, userRepo)
	jwt := _jwt.NewJwtService(userRepo)

	e := echo.New()
	g := e.Group("/api/v1")

	userJwt := e.Group("")
	jwt.SetJwtAdmin(userJwt)

	_carHandler.NewHandler(userJwt, carService)
	_garageHandler.NewHandler(g, garageService)
	_userHandler.NewHandler(g, userService)
	_orderHandler.NewHandler(g, orderService)

	PORT := os.Getenv("SERVER_PORT")
	log.Fatal(e.Start(PORT))
}
