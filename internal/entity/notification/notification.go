package entity

import "time"

type Status string

const (
	Pending Status = "´PENDING"
	Sent    Status = "SENT"
	Failed  Status = "FAILED"
)

type Notification struct {
	ID        string
	UserID    string
	Content   string
	Channel   string
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
}
