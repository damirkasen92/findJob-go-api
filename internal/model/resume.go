package model

import "time"

type Resume struct {
	ID uint

	Title  string
	About  string
	Skills string
	UserID uint

	CreatedAt time.Time
}
