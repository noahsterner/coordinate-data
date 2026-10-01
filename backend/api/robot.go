package api

import (
	"net/http"
	"encoding/json"
	
	"github.com/google/uuid"
	"backend/core"
)

func(h Handler) GetCurrentSession (w http.ResponseWriter, r *http.Request){
	var response struct {
		MapId uuid.UUID `json:"map_id"`
	}

	response.MapId = h.Backend.CurrentSession
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func(h Handler) GetRobotMode (w http.ResponseWriter, r *http.Request){
	var response struct {
		RobotMode core.Mode	`json:"robot_mode"`
	}
	
	response.RobotMode = h.Backend.Robot.Mode

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func(h Handler) SetRobotMode (w http.ResponseWriter, r *http.Request){
	var request struct {
		RobotMode core.Mode	`json:"robot_mode"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if request.RobotMode == h.Backend.Robot.Mode {
		http.Error(w, "Mode already set", http.StatusConflict)
		return
	}
		
	h.Backend.Robot.Mode = request.RobotMode
	var response struct {
		RobotMode core.Mode	`json:"robot_mode"`
		MapId uuid.UUID		`json:"map_id"`
	}

	if request.RobotMode == core.MANUAL {
		h.Backend.CurrentSession = uuid.Nil
	} else if request.RobotMode == core.AUTO {
		h.Backend.CurrentSession = uuid.New()

		if err := h.MapRepository.Create(h.Backend.CurrentSession); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	
	response.RobotMode = request.RobotMode
	response.MapId = h.Backend.CurrentSession

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
