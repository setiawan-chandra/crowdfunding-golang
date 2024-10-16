package main

import (
	"crowdfunding-golang/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// "fmt"

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/crowdfunding-golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.Name = "John Doe"
	userInput.Email = "contoh@gmail.com"
	userInput.Occupation = "Software Engineer"
	userInput.Password = "password"

	userService.RegisterUser(userInput)
}
