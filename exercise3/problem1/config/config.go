package config

import "context"

type SharedConfig struct {
	Port int    `env:"PORT"`
	Host string `env:"HOST,default=localhost"`
}

type Config struct {
	//Api *ApiConfig
	DB  *DBConfig
	Api *ApiConfig
}

func NewConfig(ctx context.Context) (*Config, error) {

	conf := &Config{}
	// DB config
	if c, err := newDBConfig(ctx); err != nil {
		return nil, err
	} else {
		conf.DB = c
	}

	// Api config
	if c, err := newApiConfig(ctx); err != nil {
		return nil, err
	} else {
		conf.Api = c
	}
	return conf, nil
}
