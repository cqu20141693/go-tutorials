package main

import (
	"github.com/cqu20141693/go-service-common/config"
	"github.com/cqu20141693/go-service-common/event"
	"github.com/cqu20141693/go-service-common/global"
	"github.com/cqu20141693/go-service-common/logger"
	"github.com/cqu20141693/go-service-common/logger/cclog"
	"github.com/cqu20141693/go-service-common/web"
	"github.com/cqu20141693/go-tutorials/registry/nacos/common"
	"github.com/cqu20141693/go-tutorials/registry/nacos/controller"
	"go.uber.org/zap/zapcore"
	"os"
)

func init() {
	global.SetLogLevel(zapcore.DebugLevel)
	config.Init()
	logger.Init()
	h := controller.Handler{}
	web.AddRouterRegister(h.InitRouterMapper)
}
func main() {
	event.TriggerEvent(event.Start)

	service := common.CreateServiceWithHttpServer()
	service.Init()
	common.ConfigRouter(service.Server())
	// Run service
	if err := service.Run(); err != nil {
		cclog.Error(err.Error())
		os.Exit(0)
	}

}
