package interfaces

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/vetzii/expo-sdk/domain"
)

const (
	HOST         = "https://exp.host"
	BASE_API_URI = "/--/api/v2"
)

type Config struct {
	AccessToken string
	ApiUri      string
	Host        string
	HttpClient  *http.Client
}

func PushNotificationClient(config *Config) *Config {
	// instance
	cC := new(Config)

	//	default
	host := HOST
	apiURL := BASE_API_URI
	httpClient := &http.Client{}
	token := ""

	//	custom setting
	if config != nil {
		if config.Host != "" {
			host = config.Host
		}
		if config.ApiUri != "" {
			apiURL = config.ApiUri
		}
		if config.AccessToken != "" {
			token = config.AccessToken
		}
		if config.HttpClient != nil {
			httpClient = config.HttpClient
		}
	}

	//	set config
	cC.Host = host
	cC.ApiUri = apiURL
	cC.HttpClient = httpClient
	cC.AccessToken = token

	return &Config{
		AccessToken: cC.AccessToken,
		ApiUri:      cC.ApiUri,
		Host:        cC.Host,
		HttpClient:  cC.HttpClient,
	}
}

//	Sends a single push notification
func (cC *Config) Send(msg *domain.PushNotificationMessage) (domain.PushNotificationResponse, error) {
	var err error
	var response []domain.PushNotificationResponse

	response, err = cC.submitProcess([]domain.PushNotificationMessage{*msg})

	if err != nil {
		return domain.PushNotificationResponse{}, err
	}

	return response[0], nil
}

//	Sends multiple push notification
func (cC *Config) SendMultiple(messages []domain.PushNotificationMessage) ([]domain.PushNotificationResponse, error) {
	return cC.submitProcess(messages)
}

func (cC *Config) submitProcess(messages []domain.PushNotificationMessage) ([]domain.PushNotificationResponse, error) {

	var jsonDTO []byte
	var err error
	var request *http.Request
	var response *http.Response
	var pushServerError domain.PushNotificationExpoServerError

	//	Validate messages list
	for _, message := range messages {

		if len(message.To) == 0 {
			return nil, errors.New("No recipients")
		}

		for _, recipient := range message.To {
			if recipient == "" {
				return nil, errors.New("Invalid push token")
			}
		}
	}
	//	set host
	uri := cC.Host + cC.ApiUri + "/push/send"

	//	json message encoding
	if jsonDTO, err = json.Marshal(messages); err != nil {
		return nil, err
	}

	//	build request
	if request, err = http.NewRequest("POST", uri, bytes.NewReader(jsonDTO)); err != nil {
		return nil, err
	}

	//	define request headers
	request.Header.Add("Content-Type", "application/json")

	//	validate access token
	if cC.AccessToken != "" {
		request.Header.Add("Authorization", "Bearer "+cC.AccessToken)
	}

	//	execute request
	if response, err = cC.HttpClient.Do(request); err != nil {
		return nil, err
	}

	//	validate request status
	if response.StatusCode >= 200 && response.StatusCode <= 299 {

		var r *domain.PushNotificationExpoResponse

		if err = json.NewDecoder(response.Body).Decode(&r); err != nil {
			return nil, err
		}

		//	error validation
		switch true {
		//	error in the request
		case r.Errors != nil:

			pushServerError = domain.PushNotificationExpoServerError{}
			pushServerError.Message = "Invalid server response"
			pushServerError.Response = response
			pushServerError.ResponseData = r
			pushServerError.Errors = r.Errors
			return nil, pushServerError

		//	Error in the response format
		case r.Data == nil:

			pushServerError = domain.PushNotificationExpoServerError{}
			pushServerError.Message = "Invalid server response"
			pushServerError.Response = response
			pushServerError.ResponseData = r
			pushServerError.Errors = r.Errors
			return nil, pushServerError

		//	error validating response integrity
		case len(messages) != len(r.Data):
			pushServerError = domain.PushNotificationExpoServerError{}
			pushServerError.Message = fmt.Sprintf("Mismatched response length. Expected %d receipts but only received %d", len(messages), len(r.Data))
			pushServerError.Response = response
			pushServerError.ResponseData = r
			pushServerError.Errors = r.Errors

			return nil, pushServerError
		}

		// Add the original message to each response for reference
		for i := range r.Data {
			r.Data[i].PushMessage = messages[i]
		}

		return r.Data, nil
	}

	return nil, fmt.Errorf("Error invalid response (%d %s)", response.StatusCode, response.Status)
}
