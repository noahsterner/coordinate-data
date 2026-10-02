package core

import (
	"sync"
)

type Mode int

const (
	MANUAL Mode = iota
	AUTO
)

type Robot struct {
	mu sync.Mutex
	mode Mode	
}

func NewRobot() *Robot {
	return &Robot{mode: MANUAL}
}

func(r *Robot) GetMode() Mode {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.mode
}

func(r *Robot) SetMode(mode Mode) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.mode = mode
}
