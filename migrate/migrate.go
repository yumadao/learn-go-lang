package main

import (
	"fmt"
	"learn-go/db"
	"learn-go/model"
)

func main() {
	dbConn := db.NewDB()
	defer db.CloseDB(dbConn)
	if err := dbConn.AutoMigrate(new(model.User), new(model.Task)); err != nil {
		fmt.Println("Error Migrated", err.Error())
	} else {
		fmt.Println("Successfully Migrated")
	}
}
