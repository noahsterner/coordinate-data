package core

import(
	"github.com/google/uuid"
)

type Backend struct {
	Robot *Robot
	CurrentSession uuid.UUID
}
