package main

import (
	"fmt"
	"github.com/cqu20141693/go-service-common/logger/cclog"
	"github.com/spf13/viper"
)

func main() {
	ReadLocalYaml()
	serial := viper.GetString("sip.serial")
	cclog.Info(fmt.Sprintf("serial=%s", serial))
}
func ReadLocalYaml() {
	// 读取本地配置
	viper.SetConfigName("bootstrap.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./resource")
	viper.AddConfigPath("/etc/resource")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			cclog.Info("Config file not found; ignore error if desired")
		} else {
			cclog.Info("Config file was found but another error was produced")
		}
	}
}
