package config

import (
	"db_cp_6_sem/pkg/logger"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	Server ServerConfig   `yaml:"server"`
	User   PostgresConfig `yaml:"postgres"`
	Empl   PostgresConfig `yaml:"emplpostgres"`
	Admin  PostgresConfig `yaml:"adminpostgres"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port" default:"8080"`
}

type PostgresConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port" default:"5432"`
	Database string `yaml:"dbname"`
}

var instance *Config
var once sync.Once

func GetConfig(logger *logger.Logger) *Config {
	once.Do(func() {
		logger.Info("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig("./config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
