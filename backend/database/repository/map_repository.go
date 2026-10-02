package repository

import (
	"fmt"
	"time"
	"database/sql"
	"github.com/google/uuid"
)

type Map struct {
	UUID uuid.UUID		`json:"uuid"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
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

func(r MapRepository) FindAll() ([]Map, error) {
	query := `SELECT
	uuid, created_at, updated_at
	FROM map
	`
		
	rows, err := r.db.Query(query)
	if err != nil {
		return []Map{}, err
	}
	defer rows.Close()
	
	fmt.Println(rows)
	var maps []Map
	for rows.Next() {
		room := Map{}
		err := rows.Scan(
			&room.UUID,
			&room.CreatedAt,
			&room.UpdatedAt,
		)
		if err != nil {
			return []Map{}, err
		}

		maps = append(maps, room)
	}

	if err := rows.Err(); err != nil {
		return []Map{}, err
	}

	return maps, nil
}
