package domain

type PushNotificationResponse struct {
	PushMessage PushNotificationMessage
	ID          string            `json:"id"`
	Status      string            `json:"status"`
	Message     string            `json:"message"`
	Details     map[string]string `json:"details"`
}
