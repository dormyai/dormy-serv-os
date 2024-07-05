package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Http struct {
		Port string
	}
	Mysql struct {
		Host     string
		Port     int
		Username string
		Password string
		Db       string
	}
	API struct {
		Key string
	}
	Log struct {
		Level string
	}
	AES struct {
		KeyPath string
	}
	Key struct {
		VegaKey []byte
	}
	Chains struct {
		Main        string
		MainWss     string
		MainChainId int64
		BSC         string
		Mumbai      string
	}
	Dormy struct {
		Dormy   string
		Manager string
		Access  string
	}
	Common struct {
		Env                 string
		ServiceName         string
		OpenExchangeRateKey string
	}
}

var appCfg = &AppConfig{}

func init() {

	env := os.Getenv("DORMY_APP_ENV")

	if env == "production" {
		viper.SetConfigFile("./config/pro.yml")
	} else {
		viper.SetConfigFile("./config/dev.yml")
	}

	viper.SetConfigType("yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("无法读取配置文件：%s", err)
	}

	err = viper.Unmarshal(appCfg)
	if err != nil {
		log.Fatalf("无法解析配置：%s", err)
	}
}

func Get() *AppConfig {
	return appCfg
}
