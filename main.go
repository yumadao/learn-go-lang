package main

import (
	"learn-go/controller"
	"learn-go/db"
	"learn-go/repository"
	"learn-go/router"
	"learn-go/usecase"
)

func main() {
	// DB
	db := db.NewDB()

	// user
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	// task
	taskRepository := repository.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskController := controller.NewTaskController(taskUsecase)

	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
