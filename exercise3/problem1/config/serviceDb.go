package config

import (
	"context"
	"strconv"
)

type DBConfig struct {
	*SharedConfig
	User     string
	Password string
	DBName   string
}

func newDBConfig(ctx context.Context) (*DBConfig, error) {
	c := &DBConfig{}
	props := ctx.Value(PropName).(map[string]string)
	c.DBName = props["DB_NAME"]
	c.User = props["DB_USER"]
	c.Password = props["DB_PASSWORD"]
	port, err := strconv.Atoi(props["DB_PORT"])
	c.SharedConfig = &SharedConfig{}
	if err == nil {
		c.SharedConfig.Port = port
	}
	c.SharedConfig.Host = props["DB_HOST"]
	return c, nil
}
