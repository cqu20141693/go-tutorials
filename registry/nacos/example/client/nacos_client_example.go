package main

import (
	"context"
	"github.com/cqu20141693/go-service-common/config"
	"github.com/cqu20141693/go-service-common/event"
	"github.com/cqu20141693/go-service-common/logger"
	"github.com/cqu20141693/go-tutorials/registry/nacos/common"
	"go-micro.dev/v4/client"
	"log"
)

func init() {
	config.Init()
	logger.Init()
}
func main() {
	event.TriggerEvent(event.Start)
	c := common.CreateClient()
	//request := c.NewRequest("go-tutorials", "/health", "")
	//response := new(map[string]interface{})
	//// only support POST
	//// call service
	//err := c.Call(context.Background(), request, response)
	//log.Printf("err:%v response:%#v\n", err, response)
	cameraId := "nyXivDbD20211015"
	endpoint := "/api/device/meta/getByDeviceKey?deviceKey=" + cameraId
	req := client.NewRequest("device-backend", endpoint, map[string]string{})
	type DeviceInfo struct {
		GroupKey    string `json:"groupKey"`
		SN          string `json:"sn"`
		DeviceKey   string `json:"deviceKey"`
		DeviceToken string `json:"deviceToken"`
	}
	type result struct {
		Code    string
		Message string
		Data    DeviceInfo
	}
	resp := result{}
	err := c.Call(context.Background(), req, &resp)
	log.Printf("err:%v response:%#v\n", err, resp)
}
