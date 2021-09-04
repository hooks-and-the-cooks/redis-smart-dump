package main

import "github.com/spf13/viper"

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/")
	viper.AddConfigPath("./resources/")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
