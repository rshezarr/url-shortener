package config

import (
	"flag"
	"github.com/spf13/viper"
	"log/slog"
	"os"
	"time"
	"url-short/internal/logging"
)

type Configuration struct {
	HTTP     HTTP `yaml:"http"`
	Database DB   `yaml:"database"`
	Logger   *slog.Logger
}

type HTTP struct {
	Host        string        `yaml:"host"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idleTimeout"`
	Port        string        `yaml:"port"`
	HeaderBytes int           `yaml:"headerBytes"`
}

type DB struct {
	Host          string `yaml:"host"`
	DbIdleTimeout int    `yaml:"dbIdleTimeout"`
	Username      string `yaml:"username"`
	Password      string `yaml:"password"`
	Driver        string `yaml:"driver"`
	Port          int    `yaml:"port"`
	DbName        string `yaml:"dbName"`
}

func initConfig() error {
	//set flag
	var configPath = flag.String("config-path", "configs/", "path to config file")

	//parse flag
	flag.Parse()

	//set config file path, name and type
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(*configPath)

	//read config by property above
	return viper.ReadInConfig()
}

func GetConfig() *Configuration {
	configs := new(Configuration)

	if err := initConfig(); err != nil {
		slog.Error("error occurred while parsing configs: %v", err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&configs); err != nil {
		slog.Error("error occurred while parsing configs: %v", err)
		os.Exit(1)
	}

	configs.Logger = logging.InitLogger()

	return configs
}
