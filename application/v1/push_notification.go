package v1

import (
	"github.com/vetzii/expo-sdk/domain"
	"github.com/vetzii/expo-sdk/interfaces"
)

type PushNotificationV1 struct {
	Add         func(notificacion *domain.PushNotificationMessage) (domain.PushNotificationResponse, error)
	AddMultiple func(notificacion []domain.PushNotificationMessage) ([]domain.PushNotificationResponse, error)
}

func PushNotification(config *interfaces.Config) PushNotificationV1 {
	client := interfaces.PushNotificationClient(config)

	return PushNotificationV1{
		Add:         client.Send,
		AddMultiple: client.SendMultiple,
	}
}
