package config

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var cfg MainConfig

func Load(filename string) (*MainConfig, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(filename)
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	// Parse config to config struct
	if err := viper.Unmarshal(&cfg, func(config *mapstructure.DecoderConfig) {
		config.TagName = "env"
	}); err != nil {
		panic(fmt.Errorf("fatal error parse config file: %w", err))
	}
	return &cfg, nil
}
