package config

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"time"
)

type Configuration struct {
	HTTP     HTTP `yaml:"http"`
	Database DB   `yaml:"database"`
}

type HTTP struct {
	Host        string        `yaml:"host"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idleTimeout"`
	Port        int           `yaml:"port"`
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

func InitConfig() error {
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

	if err := viper.Unmarshal(&configs); err != nil {
		log.Fatal(err)
	}

	return configs
}
