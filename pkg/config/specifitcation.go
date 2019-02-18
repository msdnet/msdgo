package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

//Specification config
type Specification struct {
	Port int `envconfig:"PORT"`
}

//Load Config file
func Load(configFile string) (*Specification, error) {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
	}
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "No config file found")
	}

	var config Specification
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
