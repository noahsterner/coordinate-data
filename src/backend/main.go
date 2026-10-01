package main

import (
	"net/http"
	"fmt"
	"log"
	"strconv"
	"encoding/json"

	"github.com/google/uuid"
	"xyz-robotic/src/backend/database"
	"xyz-robotic/src/backend/database/repository"
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
	
	coordinateRepository := repository.NewCoordinateRepository(db)

	http.HandleFunc("GET /api/coordinates", func (w http.ResponseWriter, r *http.Request){
		var page int
		var limit int
		var mapId uuid.UUID
		var err error

		if value := r.URL.Query().Get("mapId"); value != "" {
			mapId, err = uuid.Parse(value)
			if err != nil {
				http.Error(w, "invalid page", http.StatusBadRequest)
				return
			}
		} else {
			http.Error(w, "mapId query required", http.StatusBadRequest)
			return
		}

		if value := r.URL.Query().Get("page"); value != "" {
			page, err = strconv.Atoi(value)
			if err != nil {
				http.Error(w, "invalid page", http.StatusBadRequest)
				return
			}
		}
		
		if query := r.URL.Query()["limit"]; len(query) > 0 {
			limit, err = strconv.Atoi(query[0])
			if err != nil {
				http.Error(w, "invalid limit", http.StatusBadRequest)
				return
			}
		}
		
		room := repository.Map{UUID: mapId}
		coordinates, err := coordinateRepository.List(&room, limit, page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		fmt.Println(coordinates)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(coordinates); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("POST /api/coordinates", func (w http.ResponseWriter, r *http.Request){
		var req struct {
			MapId uuid.UUID `json:"map_id"`
			X int32 `json:"x"`
			Y int32 `json:"y"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		if req.MapId == uuid.Nil {
			http.Error(w, "map_id required", http.StatusBadRequest)
		}
		
		coordinate := repository.Coordinate{
			MapId: req.MapId,
			X: req.X,
			Y: req.Y,
		}

		err := coordinateRepository.Create(&coordinate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "applicaton/json")
		w.WriteHeader(http.StatusCreated)
	})

	if err = http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
