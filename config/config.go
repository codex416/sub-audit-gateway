package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {

	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`


	Subscription struct {
		Upstream string `yaml:"upstream"`
	} `yaml:"subscription"`


	Telegram struct {
		Enable bool   `yaml:"enable"`
		BotToken string `yaml:"bot_token"`
		ChatID string `yaml:"chat_id"`
	} `yaml:"telegram"`


	Security struct {
		MaxIPChange int `yaml:"max_ip_change"`
		BlacklistEnable bool `yaml:"blacklist_enable"`
	} `yaml:"security"`

}


func LoadConfig(path string) (*Config, error) {

	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}


	var config Config


	err = yaml.Unmarshal(data, &config)

	if err != nil {
		return nil, err
	}


	return &config, nil
}
