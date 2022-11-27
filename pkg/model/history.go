package model

import "time"

type History struct {
	Id        int64     `json:"id" db:"id"`
	Event     string    `json:"event" db:"event"`
	EventTime time.Time `json:"event_time" db:"event_time"`
	UserId    int64     `json:"user_id" db:"user_id"`
}
