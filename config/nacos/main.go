package main

import (
	"fmt"
	"github.com/asim/go-micro/plugins/config/encoder/yaml/v4"
	"github.com/asim/go-micro/plugins/config/source/nacos/v4"
	"github.com/cqu20141693/go-service-common/boot"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/reader/json"
	"go-micro.dev/v4/config/source"
)

func init() {
}

// Port number
type Port uint16
type SipConfig struct {
	Serial          string `json:"serial"`
	Realm           string `json:"realm"`
	Expires         int32
	ListenAddress   string `json:"listenAddress"`
	SipIp           string // sip 服务器ip
	SipPort         Port   // sip 服务器端口
	MediaIP         string //媒体服务器地址
	MediaPort       uint16 //媒体服务器端口
	AudioEnable     bool   //是否开启音频
	CatalogInterval int32
	AutoInvite      bool
	Network         string `json:"network"`
}

func main() {
	enc := yaml.NewEncoder()
	r := json.NewReader(reader.WithEncoder(enc))
	conf, err := config.NewConfig(config.WithReader(r))
	if err != nil {
		return
	}

	var ops []source.Option
	ops = append(ops, nacos.WithAddress([]string{"172.30.203.22:8848"}))
	ops = append(ops, nacos.WithClientConfig(constant.ClientConfig{NamespaceId: "ca1d1ded-cb0b-460c-8efa-7e665c7a34e0"}))
	ops = append(ops, nacos.WithDataId("sip-link-dev.yml"))
	ops = append(ops, nacos.WithGroup("DEFAULT_GROUP"))
	ops = append(ops, source.WithEncoder(enc))
	err = conf.Load(nacos.NewSource(ops...))
	if err != nil {
		return
	} else {
		fmt.Printf("nacos config: %v \n", conf.Map())
		get := conf.Get("sip", "serial")
		fmt.Println(string(get.Bytes()))
		sipConfig := SipConfig{}
		jsonV := conf.Get("sip")
		jsonV.Scan(&sipConfig)
		fmt.Println(sipConfig)
		//
	}

	boot.ListenSignal()
}
