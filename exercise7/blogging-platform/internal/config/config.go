package config

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// LoadConfig loads the properties from the given file into a map[string]string.
func LoadConfig(filePath string) (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open configuration file: %w", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	configMap := make(map[string]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid configuration line: %s", line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		configMap[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading configuration file: %w", err)
	}

	log.Println("Config loaded successfully.")
	return configMap, nil
}
