package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func InitConfig() {
	Config = viper.New()
	Config.SetConfigName("config")  // 配置文件名
	Config.SetConfigType("yaml")    // 配置文件类型
	Config.AddConfigPath("config/") // 配置文件路径

	if err := Config.ReadInConfig(); err != nil {
		log.Fatalf("加载配置文件失败: %v", err)
	}
}
