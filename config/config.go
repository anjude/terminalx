package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os/user"
)

var BotConf Config

type ChatGPTConfig struct {
	ApiKey     string `mapstructure:"apiKey"`
	Proxy      string `mapstructure:"proxy"`
	ChatRounds int    `mapstructure:"chatRounds"`
}

type Config struct {
	ChatGPT ChatGPTConfig `mapstructure:"chatGPTConfig"`
	Version string        `mapstructure:"version"`
}

func InitConfig() error {
	currentUser, _ := user.Current()
	configFile := currentUser.HomeDir + "/.terminalx/config.yaml"
	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("read config file err: %v\n", err)
		return err
	}

	if err := viper.Unmarshal(&BotConf); err != nil {
		fmt.Printf("unmarshal config file err: %v\n", err)
		return err
	}
	return nil
}
