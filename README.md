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
	expo_interfaces "github.com/vetzii/expo-sdk/interfaces"
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
	pushNotification = expo.PushNotification(&expo_interfaces.Config{AccessToken: "EXPO_ACCESS_TOKEN"})


	//	Send notification
	if response, err = pushNotification.Add(
		&expo_sdk_domain.PushNotificationMessage{
			To:       tokens,
			Body:     "vetzii notification",
			Data:     map[string]string{"withSome": "data"},
			Sound:    "default",
			Title:    "Push Notification",
			Priority: "default", //	"default" | "normal" | "high"
		},
	); err != nil {
		fmt.Println("error sending notification: ", err.Error())
		//	option
		//	panic(err)
	}

	fmt.Println(response.Status) // "error" | "ok"
}

```

## Expo documentation

[Push Notifications Overview](https://docs.expo.dev/push-notifications/overview/)
