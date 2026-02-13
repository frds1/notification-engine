package notification

type SendNotificationInput struct {
	UserID  string `json:"user_id" binding:"required"`
	Content string `json:"content" binding:"required"`
	Channel string `json:"channel" binding:"required,oneof=sms email push"`
}
