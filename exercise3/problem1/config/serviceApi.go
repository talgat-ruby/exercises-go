package config

import (
	"context"
	"strconv"
)

type ApiConfig struct {
	*SharedConfig
}

func newApiConfig(ctx context.Context) (*ApiConfig, error) {
	c := &ApiConfig{}
	props := ctx.Value(PropName).(map[string]string)
	port, err := strconv.Atoi(props["PORT"])
	c.SharedConfig = &SharedConfig{}
	if err == nil {
		c.SharedConfig.Port = port
	}
	c.SharedConfig.Host = props["DB_HOST"]

	return c, nil
}
