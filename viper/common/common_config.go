package main

import (
	"fmt"
	"github.com/cqu20141693/go-service-common/config"
	"github.com/cqu20141693/go-service-common/event"
	"github.com/cqu20141693/go-service-common/global"
	"github.com/cqu20141693/go-service-common/logger/cclog"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

func init() {
	global.SetLogLevel(zapcore.DebugLevel)
	config.Init()
}
func main() {
	event.TriggerEvent(event.Start)
	aKey := "cc.profiles.active"
	bKey := "cc.log.dir"
	nKey := "cc.application.name"
	sKey := "server.port"
	sprintf := fmt.Sprintf("read local over. %s =%s, %s=%s %s=%s %s=%s", aKey, viper.GetString(aKey), bKey, viper.GetString(bKey), nKey, viper.GetString(nKey), sKey, viper.GetString(sKey))
	cclog.Debug(sprintf)
}
