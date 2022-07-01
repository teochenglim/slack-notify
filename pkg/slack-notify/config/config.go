package config

import (
	"encoding/json"
	"fmt"
	// "errors"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
  SLACK_COLOR    string `mapstructure:"SLACK_COLOR"`
  SLACK_CHANNEL  string `mapstructure:"SLACK_CHANNEL"`
	SLACK_ICON     string `mapstructure:"SLACK_ICON"`
  SLACK_MARKDOWN bool   `mapstructure:"SLACK_MARKDOWN"`
  SLACK_MESSAGE  string `mapstructure:"SLACK_MESSAGE"`
  SLACK_USERNAME string `mapstructure:"SLACK_USERNAME"`
  SLACK_TITLE    string `mapstructure:"SLACK_TITLE"`
  SLACK_VERBOSE  bool   `mapstructure:"SLACK_VERBOSE"`
	SLACK_WEBHOOK  string `mapstructure:"SLACK_WEBHOOK"`
}

func (c Config) String() string {
	copy := c
	copy.SLACK_WEBHOOK = "*****"
	b, err := json.Marshal(copy)
	if err != nil {
		return fmt.Sprintf("error: %s\n", err)
	}
	return string(b)
}

func (c Config) GoString() string {
	return fmt.Sprintf("config.Config: %s\n", c)
}

func LoadConfigurations() *Config {
	viper.AutomaticEnv()

	viper.SetConfigType("yaml") // file type
	viper.SetConfigName("slack-notify") // file name

	viper.AddConfigPath("./config")
  viper.AddConfigPath(".")
  viper.AddConfigPath("$HOME")

	c := &Config{}

	// err := viper.ReadInConfig()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			// fmt.Printf("Configurations file not found\n", c)
			return c
		} else {
			// Config file was found but another error was produced
			// fmt.Printf("Configurations file parsing Error\n", c)
			return c
		}
	}

	// viper.Unmarshal Unmarshal eror
	if err := viper.Unmarshal(c, func(config *mapstructure.DecoderConfig) {
		c.SLACK_MARKDOWN = true
	}); err != nil {
		fmt.Printf("Unmarshal failed, will use default settings. Error is %v", err)
	}

	if c.SLACK_VERBOSE == true {
		fmt.Printf("Configurations are %v\n", c)
	}

	return c

}
