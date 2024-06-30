package run

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/talgat-ruby/exercises-go/exercise3/problem2/internal/models"
)

func Do(rd *models.RequestData) (statusCode int) {
	slog.Info("request data", slog.String("requesData", rd.String()))
	if !isAllowMethod(rd.Method) {
		slog.Error("method not allowed", slog.String("requesData", rd.String()))
		fmt.Println("METHOD NOT ALLOWEDs")
		return 1
	}

	req, err := http.NewRequest(rd.Method, rd.URL, bytes.NewBuffer(rd.Data))
	if err != nil {
		slog.Error("Error creating request", slog.Any("err", err), slog.String("requestData", rd.String()))
		fmt.Println("ERROR:\ncan not create request")
		return 1
	}

	for _, header := range rd.Headers {
		req.Header.Set(header[0], header[1])
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Error executing request", slog.Any("err", err), slog.String("requestData", rd.String()))
		fmt.Println("ERROR:\nerror executing request")
		return 1
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Error reading response body", slog.Any("err", err), slog.String("requestData", rd.String()))
		fmt.Println("ERROR:\nerror reading response body")
		return 1
	}

	// Print the response body
	fmt.Println(string(body))
	return 0
}

func isAllowMethod(method string) bool {
	switch method {
	case http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete:
		return true
	default:
		return false
	}
}
