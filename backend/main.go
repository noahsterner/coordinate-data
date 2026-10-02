package main

import (
	"net/http"
	"fmt"
	"log"
		
	"github.com/google/uuid"

	"backend/database"
	"backend/database/repository"

	"backend/api"
	"backend/core"
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
	http.HandleFunc("GET /api/sessions", handler.GetMaps)

	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
		}

		http.DefaultServeMux.ServeHTTP(w,r)
	}))
}
