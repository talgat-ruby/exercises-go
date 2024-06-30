package cli

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/talgat-ruby/exercises-go/exercise3/problem2/internal/models"
)

func parseFlag() (*models.RequestData, int) {
	var headers models.MultiFlag
	data := flag.String("d", "", "request data")
	help := flag.Bool("h", false, "")
	flag.Var(&headers, "H", "Header values")
	// output := flag.String("output", "", "")
	flag.Usage = func() {
		message := fmt.Sprintln("Usage: go run . [METHOD] [URL] [OPTION]")
		message = fmt.Sprintf("\n%s[METHOD] and [URL] are required\n\n", message)
		message = fmt.Sprintf("%s[METHOD] support only methods: GET, POST, PUT, PATCH, DELETE\n", message)
		message = fmt.Sprintf("%sOptions:\n-H for request headers\n-d for request data\n", message)
		message = fmt.Sprintf("%sEX: go run . GET https://pokeapi.co/api/v2/pokemon -H 'accept: application/json'", message)
		fmt.Println(message)
	}

	flag.Parse()
	if *help {
		flag.Usage()
		return nil, 0
	}
	if len(os.Args) < 3 {
		flag.Usage()
		return nil, 1
	}
	flag.CommandLine.Parse(os.Args[3:])
	slog.Info("flag args", slog.Any("args", flag.Args()))

	method := os.Args[1]
	url := os.Args[2]

	slog.Info("d", slog.String("data", *data))

	return &models.RequestData{
		Method:  method,
		URL:     url,
		Headers: formatHeaders(headers),
		Data:    []byte(*data),
	}, 0
}

func formatHeaders(headers []string) [][]string {
	formattedHeader := make([][]string, 0)
	for _, v := range headers {
		header := strings.SplitN(v, ":", 2)
		if len(header) != 2 {
			slog.Error("invalid header", slog.String("header", v))
			continue
		}

		formattedHeader = append(formattedHeader, header)
	}
	return formattedHeader
}
