package config

type Config struct {
	api *Api
}

func New() *Config {
	return &Config{api: &Api{}}
}

func (c *Config) GetApi() *Api {
	return c.api
}
