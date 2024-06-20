package config

import "github.com/spf13/viper"

type Config struct {
	Port                  string `mapstructure:"PORT"`
	ExploriteAuth         string `mapstructure:"Explorite_Auth"`
	ExploritePost         string `mapstructure:"Explorite_Post"`
	ExploriteChat         string `mapstructure:"Explorite_chat"`
	ExploriteNotification string `mapstructure:"Explorite_Notification"`
	Access_key_ID         string `mapstructure:"Access_key_ID"`
	Secret_access_key     string `mapstructure:"Secret_access_key"`
	AWSRegion             string `mapstructure:"AWSRegion"`
	KafkaPort             string `mapstructure:"KAFKA_PORT"`
	KafkaTopic            string `mapstructure:"KAFKA_TOPIC"`
}

var envs = []string{
	"PORT", "Explorite_Auth", "Explorite_Post", "Explorite_Chat", "Access_key_ID", "Secret_access_key", "AWSRegion", "KAFKA_PORT", "KAFKA_TOPIC", "Explorite_Notification",
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
