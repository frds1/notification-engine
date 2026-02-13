package notification

import (
	"net/http"
	entity "notification-engine/internal/entity/notification"
	usecase "notification-engine/internal/usecase/notification"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	SendNotificationUseCase *usecase.SendNotificationUseCase
}

func NewNotificationHandler(usecase *usecase.SendNotificationUseCase) *NotificationHandler {
	return &NotificationHandler{
		SendNotificationUseCase: usecase,
	}
}

func (h *NotificationHandler) HandleSend(c *gin.Context) {
	var (
		input SendNotificationInput
		err   error
	)

	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.SendNotificationUseCase.Create(c.Request.Context(), &entity.Notification{
		Content: input.Content,
		Channel: input.Channel,
		UserID:  input.UserID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, nil)
}
