package util

import (
  "time"

  "github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environmnet variables.
type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
  DBSource string `mapstructure:"DB_SOURCE"`
  ServerAddress string `mapstructure:"SERVER_ADDRESS"`
  TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
  AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

// LoadConfig reads confirguration from environment variables.
func LoadConfig(path string) (config Config, err error) {
  viper.AddConfigPath(path)
  viper.SetConfigName("app")
  viper.SetConfigType("env") // json or xml

  viper.AutomaticEnv()

  err = viper.ReadInConfig()
  if err != nil {
    return
  }

  err = viper.Unmarshal(&config)
  return
}