package main

import (
	"log"
	"qlambi/handler"
	"qlambi/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/qlambiDB?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	userRepository := user.UserRepository((db))
	userService := user.UserService(userRepository)

	userHandler := handler.UserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	router.Run()
}

/*
User
	0. id
	1. name
	2. email
	3. password_hash
	4. avatar_file_name
Product
Transactions
Testimonial
Portofolio
*/