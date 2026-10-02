package core

import(
	"sync"
	"github.com/google/uuid"
)

type Backend struct {
	mu sync.Mutex
	Robot *Robot
	currentMap uuid.UUID
}

func NewBackend(robot *Robot) *Backend {
	return &Backend{Robot: robot, currentMap: uuid.Nil}
}

func(b *Backend) GetCurrentMap() uuid.UUID {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.currentMap
}

func(b *Backend) SetCurrentMap(mapUUID uuid.UUID) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.currentMap = mapUUID
}
