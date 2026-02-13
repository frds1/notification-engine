package notification

import (
	"context"
	entity "notification-engine/internal/entity/notification"
)

type SendNotificationUseCase struct {
	Repo entity.INotification
}

func NewSendNotificationUseCase(repo entity.INotification) *SendNotificationUseCase {
	return &SendNotificationUseCase{Repo: repo}
}

func (u *SendNotificationUseCase) Create(ctx context.Context, input *entity.Notification) error {
	return u.Repo.Create(ctx, input)
}
