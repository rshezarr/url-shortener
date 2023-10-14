package config

import (
	"flag"
	"github.com/spf13/viper"
	"log"
)

type Configuration struct {
	HTTP     HTTP `yaml:"http"`
	Database DB   `yaml:"db"`
}

type HTTP struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	Timeout     string `yaml:"timeout"`
	IdleTimeout string `yaml:"idle_timeout"`
}

type DB struct {
	IdleTimeout string `yaml:"idle_timeout"`
	Port        int    `yaml:"port"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
}

func InitConfig() error {
	//set flag
	var configPath = flag.String("config-path", "configs/", "path to config file")

	//parse flag
	flag.Parse()

	//set config file path, name and type
	viper.SetConfigName("configs")
	viper.SetConfigType("yaml")
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
