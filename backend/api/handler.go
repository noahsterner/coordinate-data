package api

import (
	"backend/database/repository"
	"backend/core"
)

type Handler struct {
	Backend *core.Backend
	CoordinateRepository *repository.CoordinateRepository
	MapRepository *repository.MapRepository
}
