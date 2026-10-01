package main

import (
	"net/http"
	"fmt"
	"log"
		
	"github.com/google/uuid"

	"xyz-robotic/src/backend/database"
	"xyz-robotic/src/backend/database/repository"

	"xyz-robotic/src/backend/api"
	"xyz-robotic/src/backend/core"
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
	
	backend := core.Backend{
		Robot: &core.Robot{Mode: core.MANUAL},
		CurrentSession: uuid.Nil,
	}

	handler := api.Handler{
		CoordinateRepository: repository.NewCoordinateRepository(db),
		MapRepository: repository.NewMapRepository(db),
		Backend: &backend,
	}

	http.HandleFunc("GET /api/coordinates", handler.GetCoordinates)
	http.HandleFunc("POST /api/coordinates", handler.PostCoordinate)

	http.HandleFunc("POST /api/robot/mode", handler.SetRobotMode)
	http.HandleFunc("GET /api/robot/mode", handler.GetRobotMode)

	http.HandleFunc("GET /api/robot/session", handler.GetCurrentSession)

	if err = http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
