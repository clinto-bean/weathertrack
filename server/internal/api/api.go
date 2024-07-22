package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"

	alerts "github.com/clinto-bean/weathertrack/internal/alerts"
)

type Alert struct {
	AlertType string `json:"type"`
	Start     string `json:"effective"`
	End       string `json:"ends"`
}

type ApiConfig struct {
	// instance string
}

func (cfg *ApiConfig) SecureHandler(w http.ResponseWriter, r *http.Request, payload interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Extract API Key

		token := r.Header.Get("Authorization")
		if len(token) < 6 || token[:6] != "token " {
			log.Println("Invalid authorization token")
			respondWithError(w, http.StatusBadRequest, "bad token format")
			return
		}

		// Validate API Key

		/* */

		// var resource interface{}

		// switch r.URL.Path {
		// 	case "/v1/alerts":
		// 		handlerGetAllAlerts(w, r)
		// }

	}
}

// Alerts Handlers

func HandlerHandleAlertRequests(w http.ResponseWriter, r *http.Request) {

	LogRequest(r)

	query, err := parseParams(r.URL.RawQuery)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "bad request format")
		return
	}
	if query == nil {
		err := HandlerGetAllAlerts(w, r)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	for key, value := range query {
		switch key {
		case "state":
			if err := HandlerGetAlertsByState(w, value); err != nil {
				respondWithError(w, http.StatusBadRequest, err.Error())
			}
			return
		default:
			{
				respondWithJSON(w, http.StatusNotFound, "resource not found")
				return
			}
		}
	}
}

func HandlerGetAllAlerts(w http.ResponseWriter, r *http.Request) error {

	data, err := alerts.GetAllAlerts()
	if err != nil {
		return err
	}
	respondWithJSON(w, http.StatusOK, data)
	return nil
}

func HandlerGetAlertsByState(w http.ResponseWriter, state string) error {
	data, err := alerts.GetAlertsByState(state)
	log.Printf("API Request: GET area=%v\n", state)
	if err != nil {
		return err
	}
	respondWithJSON(w, http.StatusOK, data)
	return nil
}

func parseParams(rawQuery string) (map[string]string, error) {
	params, err := url.ParseQuery(rawQuery)
	if err != nil {
		return nil, errors.New("could not parse query parameters")
	}

	returnVals := make(map[string]string)

	for key, values := range params {
		if len(values) > 0 {
			returnVals[key] = values[0]
		}
	}

	if len(returnVals) == 0 {
		return nil, nil
	}

	return returnVals, nil
}

func LogRequest(r *http.Request) {
	log.Printf("Request received: %v %v", r.Method, r.URL.String())
}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(status)
	w.Write(dat)
	log.Println("RespondWithJSON: Response sent.")

}

func respondWithError(w http.ResponseWriter, status int, msg string) {
	log.Println("RespondWithError: Sending error.")
	respondWithJSON(w, status, map[string]string{"error": msg})
}
