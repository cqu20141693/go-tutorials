package controller

import (
	"github.com/cqu20141693/go-service-common/logger/cclog"
	"github.com/cqu20141693/go-service-common/web"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	web.BaseRestController
}

func (h *Handler) health(ctx *gin.Context) {
	cclog.Info("method=" + ctx.Request.Method)
	h.ResponseData(ctx, map[string]string{"status": "up"})
}

func (h *Handler) InitRouterMapper(engine *gin.Engine) {
	engine.POST("/health", h.health)
	engine.GET("/health", h.health)
}
