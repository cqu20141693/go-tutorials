package main

import (
	"github.com/cqu20141693/go-service-common/logger/cclog"
	"github.com/cqu20141693/go-service-common/web"
	"github.com/cqu20141693/go-tutorials/handler"
	"github.com/gin-gonic/gin"
)

func init() {
	web.AddRouterRegister(func(router *gin.Engine) {
		cclog.Info("init router func")
		router.POST("/health", handler.Handler.Health)
	})
}
