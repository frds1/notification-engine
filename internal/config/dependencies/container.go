package dependencies

import (
	repo "notification-engine/internal/infrastructure/repository/notification"
	"notification-engine/internal/interfaces"
	handler "notification-engine/internal/interfaces/notification"
	usecase "notification-engine/internal/usecase/notification"

	"gorm.io/gorm"
)

type Container struct {
	Server *interfaces.Server
}

func NewContainer(db *gorm.DB) *Container {
	notificationRepo := repo.NewNotificationRepository(db)

	sendUseCase := usecase.NewSendNotificationUseCase(notificationRepo)

	notificationHandler := handler.NewNotificationHandler(sendUseCase)

	server := interfaces.NewServer(notificationHandler)

	return &Container{
		Server: server,
	}
}
