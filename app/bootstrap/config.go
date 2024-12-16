package bootstrap

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	env string
}

func getConfigPath(env string) string {
	if env != "production" {
		return "configs/config-dev.yml"
	}
	return "configs/config-prod.yml"
}

func NewConfig(env string) (*Config, error) {
	viper.SetConfigName(getConfigPath(env))
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	config.env = env
	return &config, nil
}

func (c *Config) GetDBUrl() string {
	if c.env == "production" {
		return fmt.Sprintf("postgres://%s:%s@%s/%s", c.Database.Username, c.Database.Password, c.Database.Host, c.Database.Name)
	}
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.Database.Username, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.Name)
}
