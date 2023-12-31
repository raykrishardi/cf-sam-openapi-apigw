package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type JSONResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1024 * 1024 // one megabyte
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)

	dec.DisallowUnknownFields()

	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}

func WriteErrorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload JSONResponse
	payload.Error = true
	payload.Message = err.Error()

	return WriteJSON(w, statusCode, payload)
}

func GetAPIGatewayErrorResponse(status int, err error) (events.APIGatewayProxyResponse, error) {
	payload := JSONResponse{
		Error:   true,
		Message: err.Error(),
	}

	out, marshalErr := json.Marshal(payload)
	if marshalErr != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError, Body: marshalErr.Error()}, marshalErr
	}

	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(out),
	}, nil
}
