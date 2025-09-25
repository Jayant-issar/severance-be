package config

import (
	"log"

	"github.com/spf13/viper"
)

// config stores all the configuration for the application
// the values are ready by viper from a config file or enviroment variables.
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig reads configuration from tfile or enviroment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app") //name of the config file
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("Warning: could not read config file: %v. Using enviroment variables.", err)
	}

	err = viper.Unmarshal(&config)
	return

}
