package models

import "fmt"

type RequestData struct {
	Method  string
	URL     string
	Headers [][]string
	Data    []byte
}

func (rd *RequestData) String() string {
	return fmt.Sprintf("Method: '%s'\nURL: '%s'\nHeaders: '%+v'\nData '%s'\n",
		rd.Method, rd.URL, rd.Headers, rd.Data)
}
