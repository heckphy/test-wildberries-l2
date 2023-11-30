package models

import "time"

type Event struct {
	Uid         string
	Title       string
	Description string
	StartsAt    time.Time
	EndsAt      time.Time
}
