package database

import "time"

type Status string

const (
	Pending Status = "pending"
	Sent    Status = "sent"
	Failed  Status = "failed"
)

type Notification struct {
	ID        string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID    string    `gorm:"type:varchar(255);index;not null" json:"user_id"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	Channel   string    `gorm:"type:varchar(50);not null" json:"channel"`
	Status    Status    `gorm:"type:varchar(50);not null" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
