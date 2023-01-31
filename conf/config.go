package conf

import (
	"github.com/spf13/viper"
	"log"
)

func ViperEnvVariable(key string, path string) string {
	viper.SetConfigName("app")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}
