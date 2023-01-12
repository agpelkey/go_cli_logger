package models

import "time"

type Logbook struct {
	Entry string
	CreatedAt time.Time
}