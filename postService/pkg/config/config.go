package config

import "github.com/spf13/viper"

type Config struct {
	DBHost            string `mapstructure:"DB_HOST"`
	DBName            string `mapstructure:"DB_NAME"`
	DBUser            string `mapstructure:"DB_USER"`
	DBPort            string `mapstructure:"DB_PORT"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	Port              string `mapstructure:"PORT"`
	Access_key_ID     string `mapstructure:"Access_key_ID"`
	Secret_access_key string `mapstructure:"Secret_access_key"`
	AWSRegion         string `mapstructure:"AWSRegion"`
	KafkaBrokers      string `mapstructure:"KAFKA_BROKERS"`
	KafkaTopic        string `mapstructure:"KAFKA_TOPIC"`
	Explorite_Auth    string `mapstructure:"Explorite_Auth"`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", "PORT", "Access_key_ID", "Secret_access_key", "AWSRegion", "KAFKA_BROKERS", "KAFKA_TOPIC", "Explorite_Auth",
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
