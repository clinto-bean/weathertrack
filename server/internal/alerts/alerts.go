package alerts

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

const baseURL = "https://api.weather.gov/alerts"

type Alerts struct {
	Type     string `json:"type"`
	Features []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Properties struct {
			ID          string    `json:"id"`
			AreaDesc    string    `json:"areaDesc"`
			Sent        time.Time `json:"sent"`
			Effective   time.Time `json:"effective"`
			Onset       time.Time `json:"onset"`
			Expires     time.Time `json:"expires"`
			Ends        time.Time `json:"ends"`
			Status      string    `json:"status"`
			MessageType string    `json:"messageType"`
			Category    string    `json:"category"`
			Severity    string    `json:"severity"`
			Certainty   string    `json:"certainty"`
			Urgency     string    `json:"urgency"`
			Event       string    `json:"event"`
			Sender      string    `json:"sender"`
			SenderName  string    `json:"senderName"`
			Headline    string    `json:"headline"`
			Description string    `json:"description"`
			Response    string    `json:"response"`
			Parameters  struct {
				AdditionalProp1 []any `json:"additionalProp1"`
				AdditionalProp2 []any `json:"additionalProp2"`
				AdditionalProp3 []any `json:"additionalProp3"`
			} `json:"parameters"`
		} `json:"properties"`
	} `json:"features"`
	Title      string    `json:"title"`
	Updated    time.Time `json:"updated"`
	Pagination struct {
		Next string `json:"next"`
	} `json:"pagination"`
}

func GetAllAlerts() (interface{}, error) {
	var u string = fmt.Sprintf("%v/active?status=actual", baseURL)
	log.Printf("Alerts: GET %v\n", u)

	resp, err := http.Get(u)
	if err != nil {
		return nil, errors.New("could not fetch alerts")
	}
	defer resp.Body.Close()

	data, err := decode(resp)
	if err != nil {
		return nil, err
	}

	if len(data.Features) == 0 {
		return map[int]string{1: "no results found"}, nil
	}

	return data.Features, nil
}

func GetAlertsByState(state string) (interface{}, error) {
	state = strings.ToUpper(state)
	var u string = fmt.Sprintf("%v/active?area=%v", baseURL, state)
	log.Printf("Alerts: GET %v\n", u)
	resp, err := http.Get(u)
	if err != nil {
		return nil, fmt.Errorf("the parameters of \"area=%v\" are invalid", state)
	}
	defer resp.Body.Close()
	data, err := decode(resp)
	if err != nil {
		return nil, err
	}
	if len(data.Features) == 0 {
		return "no alerts found", nil
	}
	return data.Features, nil
}

func decode(resp *http.Response) (Alerts, error) {
	var data Alerts
	decoder := json.NewDecoder(resp.Body)
	err := decoder.Decode(&data)
	if err != nil {
		return Alerts{}, errors.New("could not decode response: " + err.Error())
	}
	return data, nil
}
