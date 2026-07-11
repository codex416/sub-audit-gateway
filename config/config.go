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
		Enable bool `yaml:"enable"`
		BotToken string `yaml:"bot_token"`
		ChatID string `yaml:"chat_id"`
	} `yaml:"telegram"`


	Security struct {
		MaxIPChange int `yaml:"max_ip_change"`
		BlacklistEnable bool `yaml:"blacklist_enable"`
	} `yaml:"security"`

}



func LoadConfig(path string) (*Config,error){


	data,err:=os.ReadFile(path)

	if err!=nil{

		return nil,err

	}



	var cfg Config


	err=yaml.Unmarshal(
		data,
		&cfg,
	)


	if err!=nil{

		return nil,err

	}



	// 环境变量覆盖


	if v:=os.Getenv("PORT"); v!=""{

		cfg.Server.Port=8080

	}



	if v:=os.Getenv("SUBSCRIPTION_UPSTREAM");v!=""{

		cfg.Subscription.Upstream=v

	}



	if v:=os.Getenv("TELEGRAM_BOT_TOKEN");v!=""{

		cfg.Telegram.BotToken=v

	}



	if v:=os.Getenv("TELEGRAM_CHAT_ID");v!=""{

		cfg.Telegram.ChatID=v

	}



	return &cfg,nil

}
