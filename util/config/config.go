package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var ConfigServer *viper.Viper

func InitConfig() {
	// 加载配置
	ConfigServer = viper.New()
	// 获取环境变量
	//ConfigServer.AutomaticEnv()
	ConfigServer.SetConfigName("config")
	ConfigServer.SetConfigType("toml")
	ConfigServer.AddConfigPath("./config")
	err := ConfigServer.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
		panic("Fatal error config file")
	}
	ConfigServer.WatchConfig()
	ConfigServer.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})
}
