package config

import (
	"log"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func initConf(envPath string) {
	data, err := os.ReadFile(envPath)
	if err != nil {
		log.Fatalf("bruh %v\n", err)
	}

	setEnv(data)
}

func setEnv(data []byte) {
	envs := strings.Split(string(data), "\n")
	for _, v := range envs {
		s := strings.TrimSpace(v)
		if s == "" {
			continue
		}

		env := strings.SplitN(s, "=", 2)
		if len(env) != 2 {
			slog.Error("invalid formaat of env")
			continue
		}

		err := os.Setenv(env[0], env[1])
		if err != nil {
			slog.Error("setenv error", slog.Any("err", err))
		}
	}
}

// getEnv - считать environment в формете string
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// getEnvAsInt - считать environment в формете int
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}
