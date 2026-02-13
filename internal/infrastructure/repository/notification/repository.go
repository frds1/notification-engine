package database

import (
	"context"
	"errors"
	entity "notification-engine/internal/entity/notification"

	"gorm.io/gorm"
)

// NotificationRepository implementa a interface de Domínio
type NotificationRepository struct {
	db *gorm.DB
}

// Factory recebe a conexão do GORM
func NewNotificationRepository(db *gorm.DB) entity.INotification {
	return &NotificationRepository{
		db: db,
	}
}

func (r *NotificationRepository) Create(ctx context.Context, notification *entity.Notification) error {
	model := toModel(notification)

	result := r.db.WithContext(ctx).Create(model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *NotificationRepository) GetByID(ctx context.Context, id string) (*entity.Notification, error) {
	var model Notification

	result := r.db.WithContext(ctx).First(&model, "id = ?", id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return fromModel(model), nil
}

func toModel(n *entity.Notification) *Notification {
	return &Notification{
		ID:      n.ID,
		UserID:  n.UserID,
		Content: n.Content,
		Channel: n.Channel,
		Status:  Status(n.Status),
	}
}

func fromModel(m Notification) *entity.Notification {
	return &entity.Notification{
		ID:        m.ID,
		UserID:    m.UserID,
		Content:   m.Content,
		Channel:   m.Channel,
		Status:    string(m.Status),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
