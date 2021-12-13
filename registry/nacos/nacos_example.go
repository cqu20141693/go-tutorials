package main

import (
	"github.com/asim/go-micro/plugins/registry/nacos/v4"
	httpServer "github.com/asim/go-micro/plugins/server/http/v4"
	"github.com/cqu20141693/go-service-common/config"
	"github.com/cqu20141693/go-service-common/event"
	"github.com/cqu20141693/go-service-common/global"
	"github.com/cqu20141693/go-service-common/logger"
	"github.com/cqu20141693/go-service-common/logger/cclog"
	"github.com/cqu20141693/go-service-common/web"
	"github.com/cqu20141693/go-tutorials/registry/nacos/controller"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/spf13/viper"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
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

	service := CreateServiceWithHttpServer()
	service.Init()
	configRouter(service.Server())
	// Run service
	if err := service.Run(); err != nil {
		cclog.Error(err.Error())
		os.Exit(0)
	}

}

func configRouter(server server.Server) {

	hd := server.NewHandler(web.Engine)
	if err := server.Handle(hd); err != nil {
		cclog.Error(err.Error())
	}
}

func CreateRegistry() registry.Registry {
	clientConfig := constant.ClientConfig{}
	err := viper.UnmarshalKey("cc.cloud.nacos.config", &clientConfig)
	if err != nil {
		return nil
	}
	addr := config.GetStringOrDefault("cc.cloud.nacos.server-addr", "localhost:8848")
	addrs := strings.Split(addr, ",")
	return nacos.NewRegistry(nacos.WithAddress(addrs), nacos.WithClientConfig(clientConfig))
}

func CreateServiceWithHttpServer() micro.Service {
	webAddr := config.GetStringOrDefault("server.port", "8080")
	appName := config.GetStringOrDefault("cc.application.name", "go.micro")
	srv := httpServer.NewServer(
		server.Name(appName),
		server.Address(":"+webAddr),
	)

	return micro.NewService(
		micro.Server(srv),
		micro.Name(appName),
		micro.Registry(CreateRegistry()),
	)
}
