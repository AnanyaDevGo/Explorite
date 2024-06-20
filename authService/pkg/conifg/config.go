package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port                   string `mapstructure:"PORT"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBName                 string `mapstructure:"DB_NAME"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBPassword             string `mapstructure:"DB_PASSWORD"`
	DBPort                 string `mapstructure:"DB_PORT"`
	ExploriteChat          string `mapstructure:"Explorite_chat"`
	Explorite_notification string `mapstructure:"Explorite_notification"`
	KafkaBrokers           string `mapstructure:"KAFKA_BROKERS"`
	KafkaTopic             string `mapstructure:"KAFKA_TOPIC"`
}

var envs = []string{
	"PORT", "DB_HOST", "DB_NAME", "DB_USER", "DB_PASSWORD", "DB_PORT", "Explorite_chat","Explorite_notification", "KAFKA_BROKERS", "KAFKA_TOPIC",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil

}
