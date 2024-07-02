package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/talgat-ruby/exercises-go/exercise3/problem2/utils"
	"io"
	"net/http"
	"os"
	"strings"
)

type request struct {
	method  string
	url     string
	headers map[string]string
	data    string
}

func main() {
	args := os.Args
	// if no args
	if len(args) < 2 || len(args)%2 != 0 {
		fmt.Print(utils.HELP)
		return
	}
	// if only 1 arg which is -h or website
	if len(args) == 2 {
		if args[1] == "-h" || args[1] == "--help" {
			fmt.Print(utils.HELP)
		} else {
			sendRequest("GET", args[1], "", nil)
		}
		return
	}

	args, url := extractArgsAndUrl(args[1:])
	req := request{url: url, headers: map[string]string{}}

	// parse each argument
	for i := 0; i < len(args); i += 2 {
		err := processOption(args[i], args[i+1], &req)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Print(utils.HELP)
			return
		}
	}

	if req.method == "" && req.data != "" {
		req.method = "POST"
	} else if req.method == "" {
		fmt.Println(utils.HELP)
		return
	}
	sendRequest(req.method, req.url, req.data, req.headers)
}

func extractArgsAndUrl(args []string) ([]string, string) {
	var url string
	var index int
	for i := 0; i < len(args); i += 2 {
		arg := args[i]
		if arg == "-h" || arg == "--help" || arg == "-d" || arg == "--data" ||
			arg == "-H" || arg == "--header" || arg == "-X" || arg == "--method" {
			continue
		}
		url = arg
		index = i
		break
	}
	return append(args[:index], args[index+1:]...), url
}

func processOption(opt, val string, req *request) error {
	if opt == "-h" || opt == "--help" {
		return errors.New(utils.HELP)
	} else if opt == "-d" || opt == "--data" {
		if !utils.IsJSON(val) {
			return errors.New("Data is not in JSON format: " + val)
		}
		req.data = val
	} else if opt == "-H" || opt == "--header" {
		if hds, ok := utils.ParseHeader(val); !ok {
			return errors.New("Invalid header option: " + val)
		} else {
			for k, v := range hds {
				req.headers[k] = v
			}
		}
	} else if opt == "-X" || opt == "--method" {
		if method, ok := utils.ParseMethod(val); ok {
			req.method = method
		} else {
			return errors.New("Invalid method: " + val)
		}
	} else {
		return errors.New("Invalid option: " + opt + " " + val)
	}
	return nil
}

func sendRequest(method string, url string, body string, headers map[string]string) {
	bodyReader := utils.ParseBody(body)
	// special cases when only url is given
	if strings.HasPrefix(url, "localhost") {
		url = "http://" + url
	} else if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}
	req, reqErr := http.NewRequest(method, url, bodyReader)
	if reqErr != nil {
		fmt.Println("Cannot create request: " + reqErr.Error())
		return
	}
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	client := &http.Client{}
	resp, errDo := client.Do(req)
	if errDo != nil {
		fmt.Println("Error while sending request: " + errDo.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	printResponse(resp)
}

func printResponse(resp *http.Response) {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Not readable response")
		return
	}
	fmt.Println("Status code - " + resp.Status)
	if resp.StatusCode == http.StatusOK {
		dst := &bytes.Buffer{}
		if err := json.Indent(dst, bodyBytes, "", "  "); err == nil {
			fmt.Println("Body - " + dst.String())
			return
		}
	}
	fmt.Println("Body - " + string(bodyBytes)[:len(bodyBytes)-1])
}
