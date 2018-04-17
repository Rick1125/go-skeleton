package main

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	Health string
	DSN    string
	Redis  map[string]string
}

func NewConfig(name, file string) (*Config, error) {
	if err := readConfigs(name, file); err != nil {
		return nil,
			errors.Wrap(err, "init config file error")
	}
	return &Config{
		Health: viper.GetString("health"),
		DSN:    viper.GetString("dsn"),
		Redis:  viper.GetStringMapString("redis"),
	}, nil
}

func readConfigs(name, file string) error {
	i := strings.LastIndex(file, "/") + 1
	viper.SetConfigType("yaml")
	if i < 1 {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		//viper.AddConfigPath("$HOME/." + name)
		viper.AddConfigPath("/etc/" + name)
	} else {
		j := strings.LastIndex(file, ".")
		viper.AddConfigPath(file[0:i])
		viper.SetConfigName(file[i:j])
	}
	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "read config fail")
	}

	return nil
}
