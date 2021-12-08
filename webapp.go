package main

import (
	"github.com/cqu20141693/go-service-common/boot"
	"github.com/cqu20141693/go-service-common/logger/cclog"
	"go.uber.org/zap/zapcore"
	"os"
)

func init() {
	cclog.SetLevel(zapcore.DebugLevel)
}
func main() {

	boot.Boot(os.Args)
}
