package api

import (
	"net/http"
	"encoding/json"
	
	"github.com/google/uuid"
	"xyz-robotic/src/backend/core"
)

func(h Handler) GetCurrentSession (w http.ResponseWriter, r *http.Request){
	var response struct {
		CurrentSession uuid.UUID `json:"current_session"`
	}

	response.CurrentSession = h.Backend.CurrentSession
	
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
		response.MapId = uuid.Nil
		response.RobotMode = request.RobotMode
	} else if request.RobotMode == core.AUTO {
		h.Backend.CurrentSession = uuid.New()
		response.MapId = h.Backend.CurrentSession
		response.RobotMode = request.RobotMode
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
