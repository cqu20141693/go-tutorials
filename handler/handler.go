package handler

import (
	"github.com/cqu20141693/go-service-common/web"
	"github.com/gin-gonic/gin"
)

var Handler = handler{}

type handler struct {
	web.BaseRestController
}

func (r *handler) Health(e *gin.Context) {
	r.ResponseData(e, map[string]string{"status": "up"})
}
