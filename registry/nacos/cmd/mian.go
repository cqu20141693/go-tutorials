package main

import (
	"github.com/cqu20141693/go-service-common/boot"
	"github.com/cqu20141693/go-service-common/global"
	"github.com/cqu20141693/go-service-common/web"
	"github.com/cqu20141693/go-tutorials/registry/nacos/controller"
	"go.uber.org/zap/zapcore"
	"os"
)

func init() {
	global.SetLogLevel(zapcore.DebugLevel)
	h := controller.Handler{}
	web.AddRouterRegister(h.InitRouterMapper)
}
func main() {
	boot.Micro(os.Args)

}
