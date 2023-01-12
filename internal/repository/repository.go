package repository

import (
	"database/sql"

	"github.com/agpelkey/cli_logger/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	Add(entry models.Logbook) (string, error)
	Delete(id int) error
	FetchLogs() ([]*models.Logbook, error)
}
