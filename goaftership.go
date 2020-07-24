package goaftership

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var baseUrl string

func init() {
	baseUrl = "https://api.aftership.com/v4"
}

func post(apiKey string, uri string, requestBody []byte) (ResponseEnvelope, error) {
	client := &http.Client{}

	req, _ := http.NewRequest("POST", baseUrl+uri, bytes.NewBuffer(requestBody))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("aftership-api-key", apiKey)

	res, _ := client.Do(req)

	envelope := ResponseEnvelope{}

	err := json.NewDecoder(res.Body).Decode(&envelope)
	if err != nil {
		return ResponseEnvelope{}, err
	}

	if envelope.Meta.Code > 299 {
		return envelope, errors.New(fmt.Sprintf("AfterShip API Error [%d]: %s", envelope.Meta.Code, envelope.Meta.Message))
	}

	return envelope, nil

}

func PostTracking(apiKey string, courierSlug string, trackingNumber string) error {

	requestBody, _ := json.Marshal(&RequestEnvelope{Tracking: TrackingRequest{TrackingNumber: trackingNumber}}) // sorry garbage collector

	// don't even need the envelope (for now)
	_, err := post(apiKey, "/trackings", requestBody)

	if err != nil {
		return err
	}

	return nil

}

func GetTracking(apiKey string, courierSlug string, trackingNumber string) (*Tracking, error) {

	client := &http.Client{}

	req, _ := http.NewRequest("GET", baseUrl+"/trackings/"+courierSlug+"/"+trackingNumber, nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("aftership-api-key", apiKey)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	envelope := ResponseEnvelope{}

	err = json.NewDecoder(res.Body).Decode(&envelope)
	if err != nil {
		return nil, err
	}

	if envelope.Meta.Code > 299 {
		return &envelope.Data.Tracking, errors.New(fmt.Sprintf("AfterShip API Error [%d]: %s", envelope.Meta.Code, envelope.Meta.Message))
	}

	return &envelope.Data.Tracking, nil

}

func PostNotification(apiKey string, courierSlug string, trackingNumber string, emails []string, phones []string) error {

	requestBody, _ := json.Marshal(&RequestEnvelope{Notification: NotificationRequest{Emails: emails, SMSes: phones}}) // sorry garbage collector

	// don't even need the envelope (for now)
	_, err := post(apiKey, "/notifications/"+courierSlug+"/"+trackingNumber+"/add", requestBody)

	if err != nil {
		return err
	}

	return nil

}
