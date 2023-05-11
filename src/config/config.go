package config

import (
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"

	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
	}
	Server struct {
		Host string
		Port string
	}
	Auth struct {
		Identifier string
		Domain     string
	}
}

var C config

func ReadConfig(path string) (err error) {
	viper.AddConfigPath(path)
	//viper.SetConfigType("env")
	viper.SetConfigFile("config.yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&C); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	spew.Dump(C)
	return err
}
