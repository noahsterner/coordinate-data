package api

import (
	"net/http"
	"encoding/json"
)

func(h Handler) GetMaps (w http.ResponseWriter, r *http.Request){
	maps, err := h.MapRepository.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(maps); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
