package api

import (
	"xyz-robotic/src/backend/database/repository"
	"xyz-robotic/src/backend/core"
)

type Handler struct {
	Backend *core.Backend
	CoordinateRepository *repository.CoordinateRepository
	MapRepository *repository.MapRepository
}
