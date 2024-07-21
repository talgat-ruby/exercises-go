package request

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// MakeRequest создает и выполняет HTTP запрос
func MakeRequest(method, url, headers, data string) error {
	var req *http.Request
	var err error

	if data != "" {
		req, err = http.NewRequest(method, url, bytes.NewBuffer([]byte(data)))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return fmt.Errorf("ошибка при создании запроса: %w", err)
	}

	// Установка заголовков, если они заданы
	if headers != "" {
		hdrs := strings.Split(headers, ";")
		for _, h := range hdrs {
			parts := strings.Split(h, ":")
			if len(parts) == 2 {
				req.Header.Set(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
			}
		}
	}

	// Выполнение HTTP запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}
	defer resp.Body.Close()

	// Чтение тела ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("ошибка при чтении ответа: %w", err)
	}

	// Вывод кода статуса и тела ответа
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Println(string(body))
	} else {
		fmt.Println("Ошибка:", resp.Status)
	}

	return nil
}
