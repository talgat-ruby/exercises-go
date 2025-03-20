package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	statuserror "tracker/utils/status_error"
)

func RequestJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	ct := r.Header.Get("Content-Type")
	if ct != "" {
		typeC := strings.ToLower(ct)
		cType := strings.TrimSpace(strings.Split(typeC, ";")[0])
		if cType != "application/json" {
			msg := "Content-Type header is not application/json"
			return statuserror.New(http.StatusUnsupportedMediaType, msg)
		}
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarchalTypeError *json.UnmarshalTypeError
		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed json (at position %d)", syntaxError.Offset)
			return statuserror.New(http.StatusBadRequest, msg)
		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := "Request body contains badly-formed json"
			return statuserror.New(http.StatusBadRequest, msg)
		case errors.As(err, &unmarchalTypeError):
			msg := fmt.Sprintf(
				"Request body contains an invalid value for the %q field (at position %d)",
				unmarchalTypeError.Field,
				unmarchalTypeError.Offset,
			)
			return statuserror.New(http.StatusBadRequest, msg)
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return statuserror.New(http.StatusBadRequest, msg)
		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return statuserror.New(http.StatusBadRequest, msg)
		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			return statuserror.New(http.StatusRequestEntityTooLarge, msg)
		default:
			return err
		}

	}
	err = decoder.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		msg := "Request body must only contain a single json object"
		return statuserror.New(http.StatusBadRequest, msg)

	}
	return nil
}
