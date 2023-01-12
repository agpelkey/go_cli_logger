package logger

import (
	"database/sql"
	"fmt"
	"time"
)

// Log struct represents log entry items
type log struct {
	Entry     string
	CreatedAt time.Time
	DB *sql.DB
}

// Temporary List to hold entry items *Will update this portion to sqlite later*
type List []log




