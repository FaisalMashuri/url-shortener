package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppConfig struct {
		Name   string `mapstructure:"name"`
		Host   string `mapstructure:"host"`
		Port   string `mapstructure:"port"`
		Secret string `mapstructure:"secret"`
	} `mapstructure:"Appconfig"`

	Database struct {
		Driver   string `mapstructure:"driver"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Database string `mapstructure:"database"`
	} `mapstructure:"database"`

	Log struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Index    string `mapstructure:"index"`
	} `mapstructure:"log"`

	ErrorContract struct {
		JSONPathFile string `mapstructure:"json_path_file"`
	} `mapstructure:"errorContract"`

	Kafka struct {
		Host    string `mapstructure:"host"`
		GroupId string `mapstructure:"group_id"`
		Reset   string `mapstructure:"reset"`
	}
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read configuration file: %v", err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatalf("Failed to unmarshal configuration: %v", err)
	}
}
