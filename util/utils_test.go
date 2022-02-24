package util

import (
	"fmt"
	"go.uber.org/zap"
	"testing"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewDevelopmentConfig().Build()
}

func TestIps(t *testing.T) {
	ips, err := Ips()
	if err != nil {
		logger.Error("Ips error")
		return
	}
	logger.Info(fmt.Sprintf("ips=%s", ToJSONString(ips)))
	ip, err := GetOutBoundIP()
	if err != nil {
		logger.Error("GetOutBoundIP error")
		return
	}
	logger.Info(fmt.Sprintf("ip=%s", ip))
}
