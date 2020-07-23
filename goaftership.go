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

func PostTracking(apiKey string, courierSlug string, trackingNumber string) error {

	client := &http.Client{}

	requestBody, _ := json.Marshal(&RequestEnvelope{TrackingRequest{TrackingNumber: trackingNumber}}) // sorry garbage collector

	req, _ := http.NewRequest("POST", baseUrl+"/trackings", bytes.NewBuffer(requestBody))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("aftership-api-key", apiKey)

	res, _ := client.Do(req)

	envelope := ResponseEnvelope{}

	err := json.NewDecoder(res.Body).Decode(&envelope)
	if err != nil {
		return err
	}

	if envelope.Meta.Code > 299 {
		return errors.New(fmt.Sprintf("AfterShip API Error [%d]: %s", envelope.Meta.Code, envelope.Meta.Message))
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
