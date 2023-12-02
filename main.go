package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var config ConfigInfo

func main() {
	if err := loadConfig(); err != nil {
		fmt.Println("Error loading config: ", err)
		os.Exit(1)
	}
	runApi()
}

func loadConfig() error {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return err
	}

	return nil
}
