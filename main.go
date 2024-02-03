package main

import (
	"learn-go/controller"
	"learn-go/db"
	"learn-go/repository"
	"learn-go/router"
	"learn-go/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
