package repository

import (
	"time"

	"database/sql"
	"github.com/google/uuid"
)

type Coordinate struct {
	MapId uuid.UUID		`json:"map_id"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	X float64		`json:"x"`
	Y float64		`json:"y"`
}

type CoordinateRepository struct {
	db *sql.DB
}

func NewCoordinateRepository(db *sql.DB) *CoordinateRepository {
	return &CoordinateRepository{db: db}
}

func (r *CoordinateRepository) Create(coordinate *Coordinate) error {
	query := `INSERT INTO 
	coordinates(x, y, created_at, updated_at, map_uuid) 
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
	var err error
	var query string
	var rows *sql.Rows
	
	if(limit > 0) {
		query = `SELECT 
		x, y, map_uuid, created_at, updated_at
		FROM coordinates WHERE map_uuid = ?
		ORDER BY created_at ASC
		LIMIT ?
		OFFSET ?
		`

		rows, err = r.db.Query(query, roomMap.UUID, limit, offset)
		if err != nil {
			return nil, err
		}
	} else {
		query = `SELECT 
		x, y, map_uuid, created_at, updated_at
		FROM coordinates WHERE map_uuid = ?
		ORDER BY created_at ASC
		`

		rows, err = r.db.Query(query, roomMap.UUID)
		if err != nil {
			return nil, err
		}
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
