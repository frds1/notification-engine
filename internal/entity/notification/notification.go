package entity

import "time"

type Notification struct {
	ID        string
	UserID    string
	Content   string
	Channel   string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
