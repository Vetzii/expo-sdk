package domain

import "net/http"

type PushNotificationExpoServerError struct {
	Message      string
	Response     *http.Response
	ResponseData *PushNotificationExpoResponse
	Errors       []map[string]string
}

func (PushNotificationExpoServerError) Error() string {
	panic("unimplemented")
}
