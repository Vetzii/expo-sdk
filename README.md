# expo-sdk

SDK expo applications using go

## Installation

```
go get github.com/vetzii/expo-sdk

```

## Usage

```go
package main

import (
	"fmt"

	expo "github.com/vetzii/expo-sdk/application/v1"
	expo_sdk_domain "github.com/vetzii/expo-sdk/domain"
)

func main() {

	var response expo_sdk_domain.PushNotificationResponse
	var err error
	var tokens []string
	var pushNotification expo.PushNotificationV1

	//	tokens
	tokens = []string{}
	tokens = append(tokens, "ExponentPushToken[xxxxxxxxxxxxxxxxxxxxxx]")

	//	Instance push notification service
	pushNotification = expo.PushNotification(nil)

	// Send notification
	if response, err = pushNotification.Add(
		&expo_sdk_domain.PushNotificationMessage{
			To:       tokens,
			Body:     "Expo SDK",
			Data:     map[string]string{"withSome": "data"},
			Sound:    "default",
			Title:    "Push Notification",
			Priority: "default", //	"default" | "normal" | "high"
		},
	); err != nil {
		panic(err)
	}

	fmt.Println(response.Status) // "error" | "ok"
}

```
