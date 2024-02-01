package main

import (
	"fmt"
	"learn-go/db"
	"learn-go/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(new(model.User), new(model.Task))
}
