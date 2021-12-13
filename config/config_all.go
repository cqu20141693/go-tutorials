package main

import (
	"fmt"
	"github.com/asim/go-micro/plugins/config/encoder/yaml/v4"
	"github.com/asim/go-micro/plugins/config/source/nacos/v4"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/reader/json"
	"go-micro.dev/v4/config/source"
	"go-micro.dev/v4/config/source/file"
	"log"
)

var encoder = yaml.NewEncoder()

func main() {

	// 定义yaml解析器

	r := json.NewReader(reader.WithEncoder(encoder))
	conf, err := config.NewConfig(config.WithReader(r))
	if err != nil {
		log.Fatal(err)
	}

	loadLocalFile("./resource/bootstrap.yaml", conf)

	fmt.Println("localMap=", conf.Map())
	active := conf.Get("cc", "profiles", "active").String("")
	if active != "" {
		loadLocalFile("./resource/bootstrap-"+active+".yaml", conf)
	}
	fmt.Println("localActiveMap=", conf.Map())
	loadNacosConfig(conf, active)
	fmt.Println("nacosMap=", conf.Map())

}

func loadNacosConfig(conf config.Config, active string) {
	appName := conf.Get("cc", "application", "name").String("")

	if appName != "" {
		addr := conf.Get("cc", "cloud", "nacos", "address").StringSlice([]string{"localhost:8848"})
		clientConfig := getClientConfig(conf)
		loadAppConfig(addr, clientConfig, getDataId(conf, appName, active), getGroup(conf), conf)
		fmt.Println("nacosAppMap=", conf.Map())
		loadExtConfig(conf, addr, clientConfig)

	}
}

func loadExtConfig(conf config.Config, addr []string, clientConfig constant.ClientConfig) {
	configs := conf.Get("cc", "cloud", "nacos", "extend-configs")
	type ExtConfig struct {
		DataId  string
		Group   string
		Refresh bool
	}
	var extConfigs []ExtConfig
	err := configs.Scan(&extConfigs)
	if err != nil {
		log.Println(err)
	}
	if extConfigs != nil && len(extConfigs) > 0 {
		for _, extConfig := range extConfigs {
			if extConfig.Group != "" {
				loadAppConfig(addr, clientConfig, extConfig.DataId, extConfig.Group, conf)
			} else {
				loadAppConfig(addr, clientConfig, extConfig.DataId, "DEFAULT_GROUP", conf)
			}
		}
	}
}

func loadAppConfig(addr []string, cc constant.ClientConfig, dataId, group string, conf config.Config) {
	var ops []source.Option
	ops = append(ops, nacos.WithAddress(addr))
	ops = append(ops, nacos.WithClientConfig(cc))
	ops = append(ops, nacos.WithDataId(dataId))
	ops = append(ops, nacos.WithGroup(group))
	ops = append(ops, source.WithEncoder(encoder))
	err := conf.Load(nacos.NewSource(ops...))
	if err != nil {
		log.Printf("load nacos config failed,dataId=%s err=%s\n", dataId, err.Error())
	}
}

func getGroup(conf config.Config) string {

	return conf.Get("cc", "cloud", "nacos", "group").String("DEFAULT_GROUP")
}

func getDataId(conf config.Config, appName string, active string) string {
	fileExt := conf.Get("cc", "cloud", "nacos", "file-extension").String("yaml")
	dataId := appName + "-" + active + "." + fileExt
	return dataId
}

func getClientConfig(conf config.Config) constant.ClientConfig {
	clientConfig := constant.ClientConfig{}
	configValue := conf.Get("cc", "cloud", "nacos", "config")
	if configValue.StringMap(nil) != nil {
		_ = configValue.Scan(&clientConfig)
	}
	return clientConfig
}

func loadLocalFile(path string, conf config.Config) {
	fileSource := file.NewSource(file.WithPath(path))
	err := conf.Load(fileSource)
	if err != nil {
		log.Println("load file failed", err)
		return
	}
}
