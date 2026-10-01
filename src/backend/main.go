package main

import (
	"fmt"
	"log"
	"xyz-robotic/src/backend/database"
)

func main() {
	db, err := database.NewDatabase("./app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Database configured")

	if err := database.InitializeSchema(db); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Initialized DB Schemas")
}
