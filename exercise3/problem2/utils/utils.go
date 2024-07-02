package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

const HELP = `Usage: [OPTIONS...] <url>
	OPTIONS should be in the form of "<key> <value>"
	-h, --help info (can be used as standalone option without value)
	-H, --header "header: value",
	-d, --data '{"key1": "value1", "key2": "value2"}',
	-X, --method GET, POST, PUT, PATCH, DELETE, OPTIONS
`

var allowedMethodsFilter = map[string]bool{
	"GET":    true,
	"POST":   true,
	"PUT":    true,
	"PATCH":  true,
	"DELETE": true,
}

func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func ParseHeader(arg string) (map[string]string, bool) {
	stringArray := strings.Split(arg, ": ")
	out := make(map[string]string)
	if len(stringArray) == 0 || len(stringArray)%2 != 0 {
		return out, false
	}
	for i := 0; i < len(stringArray); i += 2 {
		out[stringArray[i]] = stringArray[i+1]
	}
	return out, true
}

func ParseBody(body string) io.Reader {
	if body == "" {
		return nil
	}
	return bytes.NewBufferString(body)
}

func ParseMethod(method string) (string, bool) {
	method = strings.ToUpper(method)
	_, ok := allowedMethodsFilter[method]
	return method, ok
}
