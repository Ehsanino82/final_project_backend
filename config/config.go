package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Database struct {
	Name     string
	Username string
	Password string
	Port     string
	Host     string
}

type Hotel struct {
	DB Database
}

func (c *Hotel) loadConfig() {
	viper.SetDefault("DB.NAME", "postgres")
	viper.SetDefault("DB.USERNAME", "postgres")
	viper.SetDefault("DB.PASSWORD", "postgres")
	viper.SetDefault("DB.HOST", "localhost")
	viper.SetDefault("DB.PORT", "5432")
}

func Load() (*Hotel, error) {
	var config Hotel

	config.loadConfig()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()

	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unmarshal config %w", err)
	}

	return &config, nil
}
