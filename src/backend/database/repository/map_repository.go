package repository

import (
	"time"
	"database/sql"
	"github.com/google/uuid"
)

type Map struct {
	UUID uuid.UUID
}

type MapRepository struct {
	db *sql.DB
}

func NewMapRepository(db *sql.DB) *MapRepository {
	return &MapRepository{db: db}
}

func(r MapRepository) Create(mapId uuid.UUID) error {
	query := `INSERT INTO 
	map(uuid, created_at, updated_at) 
	VALUES(?,?,?)
	`
		
	now := time.Now()
	result, err := r.db.Exec(query,
		mapId,
		now,
		now,
	)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}
