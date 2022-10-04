package domain

type PushNotificationExpoResponse struct {
	Data   []PushNotificationResponse `json:"data"`
	Errors []map[string]string        `json:"errors"`
}
