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
	user := user.User{
		Name: "John Doe",
	}

	userRepository.Save(user)
}
