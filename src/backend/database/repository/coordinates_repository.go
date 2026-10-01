package repository

import (
	"time"

	"database/sql"
	"github.com/google/uuid"
)

type Coordinate struct {
	MapId uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	X int32
	Y int32
}

type CoordinateRepository struct {
	db *sql.DB
}

func NewCoordinateRepository(db *sql.DB) *CoordinateRepository {
	return &CoordinateRepository{db: db}
}

func (r *CoordinateRepository) Create(coordinate *Coordinate) error {
	query := `INSERT INTO 
	coordinates(x, y, created_at, updated_at, map_id) 
	VALUES(?,?,?,?,?)
	`
		
	now := time.Now()
	result, err := r.db.Exec(query,
		coordinate.X,
		coordinate.Y,
		now,
		now,
		coordinate.MapId,
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

func (r *CoordinateRepository) List(roomMap *Map, limit, offset int) ([]Coordinate, error) {
	query := `SELECT 
	x, y, map_id, created_at, updated_at
	FROM coordinates WHERE map_id = ?
	ORDER BY created_at ASC
	LIMIT ?
	OFFSET ?
	`

	rows, err := r.db.Query(query, roomMap.UUID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var coordinates []Coordinate
	for rows.Next() {
		coordinate := Coordinate{}
		err := rows.Scan(
			&coordinate.X,
			&coordinate.Y,
			&coordinate.MapId,
			&coordinate.CreatedAt,
			&coordinate.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		coordinates = append(coordinates, coordinate)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return coordinates, nil
}
