package cmd

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

// curlCmd represents the curl command
var curlCmd = &cobra.Command{
	Use:   "http [method] [url]",
	Short: "Make HTTP requests like curl",
	Long: `A simple CLI tool to make HTTP requests. Supports methods: GET, POST, PUT, PATCH, DELETE.
Example:
	$ go run main.go http GET https://pokeapi.co/api/v2/pokemon`,
	Args: cobra.ExactArgs(2),
	Run:  runHttp,
}

var (
	headers string
	data    string
)

func init() {
	rootCmd.AddCommand(curlCmd)

	curlCmd.Flags().StringVarP(&headers, "header", "H", "", "Add HTTP headers to request")
	curlCmd.Flags().StringVarP(&data, "data", "d", "", "Data to send in request body")
}

func runHttp(cmd *cobra.Command, args []string) {
	method := strings.ToUpper(args[0])
	url := args[1]

	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	if headers != "" {
		headerPairs := strings.Split(headers, ";")
		for _, pair := range headerPairs {
			header := strings.Split(pair, ":")
			if len(header) != 2 {
				log.Fatalf("Invalid header format")
			}
			req.Header.Set(strings.TrimSpace(header[0]), strings.TrimSpace(header[1]))
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	if resp.StatusCode >= 400 {
		log.Printf("%s", body)
		fmt.Printf("%s\n", http.StatusText(resp.StatusCode))
	} else {
		fmt.Printf("%s\n", body)
	}
}
