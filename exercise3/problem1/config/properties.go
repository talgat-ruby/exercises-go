package config

import (
	"bufio"
	"log/slog"
	"os"
	"strings"
)

const PropName = "props"

func ScanDefaultProps() (map[string]string, error) {
	return ScanProps(".env")
}

func ScanProps(filename string) (map[string]string, error) {
	props := make(map[string]string)
	if len(filename) == 0 {
		filename = ".env"
	}
	slog.Info("Reading properties file")
	file, err := os.Open(filename)
	if err != nil {
		slog.Error("Failed to open properties file")
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			slog.Error("Failed to close properties file")
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		equalSignPos := strings.Index(line, "=")
		if equalSignPos < 0 {
			continue
		}
		key := strings.TrimSpace(line[:equalSignPos])
		if len(key) <= 0 || len(line) <= equalSignPos+1 {
			continue
		}
		value := strings.TrimSpace(line[equalSignPos+1:])
		if value[:1] == "\"" && value[len(value)-1:] == "\"" {
			value = value[1 : len(value)-1]
		}
		props[key] = value
	}

	if err := scanner.Err(); err != nil {
		slog.Error("Failed to read properties file")
		panic(err)
	}

	slog.Info("Finished reading properties file")
	return props, nil
}
