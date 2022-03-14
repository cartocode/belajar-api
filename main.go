package main

import (
	"log"
	"startup-api/handler"
	"startup-api/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/startupbwa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	// userByEmail, err := userRepository.FindByEmail("jsonResponse@gmail.com")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(userByEmail.Name)
	// if (userByEmail.ID == 0) {
	// 	fmt.Println("User tidak ditemukan")
	// } else {
	// 	fmt.Println(userByEmail.Name)
	// }
	// input := user.LoginInput {
	// 	Email: "jsonResponse@gmail.com",
	// 	Password: "password",
	// }
	// user, err := userService.Login(input)
	// if err != nil {
	// 	fmt.Println("Terjadi kesalahan")
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(user.Email)
	// fmt.Println(user.Name)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	
	router.Run()

	// input dari user
	// handler mapping input dari user -> ke struct
	// service mapping dari struct input ke User
	// repository save struct User to db
	// db
}
