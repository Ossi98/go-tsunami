package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewConfig(env string) *viper.Viper {
	v := viper.New()
	v.SetConfigType("yaml")                     // REQUIRED if the config file does not have the extension in the name
	v.SetConfigName(fmt.Sprintf("env.%s", env)) // name of config file (without extension)
	v.AddConfigPath(".")                        // optionally look for config in the working directory
	err := v.ReadInConfig()                     // Find and read the config file
	if err != nil {                             // Handle errors reading the config file
		log.Fatalf("fatal error config file: %v", err)
		return nil
	}

	return v
}
