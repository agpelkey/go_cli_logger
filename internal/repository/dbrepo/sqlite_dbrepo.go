package dbrepo

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/agpelkey/cli_logger/internal/models"
)

// holds our connection to the database
type SqliteDBRepo struct {
	DB *sql.DB
}

// amount of time to interact with database before request is cancelled
const dbTimeout = time.Second * 3

func (s *SqliteDBRepo) Connection() *sql.DB {
	return s.DB
}

// Add creates a new log entry to the logbook
func (s *SqliteDBRepo) Add(entry models.Logbook) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `INSERT INTO logbook (entry, completed_at) VALUES ($1, $2)`

	var newEntry string

	err := s.DB.QueryRowContext(ctx, stmt,
		entry.Entry,
		entry.CreatedAt,
	).Scan(&newEntry)

	if err != nil {
		return "", err
	}

	return newEntry, nil

}

func (s *SqliteDBRepo) FetchLogs() ([]*models.Logbook, error) {
	query, err := s.DB.Query("SELECT * FROM logbook")
	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	var entries []*models.Logbook

	for query.Next() {
		var entry models.Logbook
		err := query.Scan(
			&entry.Entry,
			&entry.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		entries = append(entries, &entry)
	}

	return entries, nil
}

// Delete method deletes a log entry
func (s *SqliteDBRepo) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `DELETE FROM logbook WHERE id = $1`

	_, err := s.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}
