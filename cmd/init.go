package cmd

import (
	"caipiao/pkg/config"

	log "github.com/sirupsen/logrus"
)

var (
	globalConfig *config.Specification
)

func initConfig() {
	var err error
	globalConfig, err = config.Load(configFile)
	if err != nil {
		log.WithError(err).Info("Could not load config from file")
	}
	log.Println("config:", globalConfig.Port)
}
