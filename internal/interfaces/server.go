package interfaces

import (
	"notification-engine/internal/interfaces/notification"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
}

func NewServer(notificationHandler *notification.NotificationHandler) *Server {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	v1 := router.Group("/api/v1")
	{
		v1.POST("/notifications", notificationHandler.HandleSend)
	}

	return &Server{Router: router}
}
