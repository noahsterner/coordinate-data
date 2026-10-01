package core

import (
	"sync"
	"github.com/google/uuid"
)

type Mode int

const (
	MANUAL Mode = iota
	AUTO
)

type Robot struct {
	mu sync.Mutex
	mode Mode	
	mapId uuid.UUID	
}

func NewRobot() *Robot {
	return &Robot{mode: MANUAL, mapId: uuid.Nil}
}

func(r *Robot) GetMode() Mode {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.mode
}

func(r *Robot) GetMapId() uuid.UUID {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.mapId
}

func(r *Robot) SetMode(mode Mode) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.mode = mode
}
func(r *Robot) SetMapId(mapId uuid.UUID) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.mapId = mapId
}
