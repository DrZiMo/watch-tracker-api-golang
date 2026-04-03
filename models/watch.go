package models

import (
	"time"

	"github.com/DrZiMo/watch-tracker-api-golang/db"
)

type Status string

const (
	StatusComplete    Status = "COMPLETED"
	StatusWatching    Status = "WATCHING"
	StatusDropped     Status = "DROPPED"
	StatusWantToWatch Status = "WANT_TO_WATCH"
)

type Watch struct {
	ID     int64
	Name   string
	Status Status
	Rating float64
	UserId int64

	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetWatches() ([]Watch, error) {
	query := `SELECT * FROM watches`
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var watches []Watch
	for rows.Next() {
		var watch Watch
		err := rows.Scan(&watch.ID, &watch.Name, &watch.Status, &watch.Rating, &watch.UserId, &watch.CreatedAt, &watch.UpdatedAt)

		if err != nil {
			return nil, err
		}

		watches = append(watches, watch)
	}

	return watches, nil
}

func (w *Watch) Create(userId int64) error {
	query := `
	INSERT INTO watches (name, status, rating, user_id)
	VALUES (?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(w.Name, w.Status, w.Rating, userId)

	if err != nil {
		return err
	}

	watchId, err := result.LastInsertId()

	if err != nil {
		return err
	}

	w.ID = watchId
	w.UserId = userId

	return nil
}
