package main

import (
	"bytes"
	"fmt"
	"github.com/asim/go-micro/plugins/config/encoder/yaml/v4"
	"github.com/cqu20141693/go-service-common/boot"
	"github.com/spf13/viper"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/reader/json"
	"go-micro.dev/v4/config/source"
	"go-micro.dev/v4/config/source/file"
)

func main() {

	encoder := yaml.NewEncoder()
	conf, err := config.NewConfig(config.WithReader(json.NewReader(reader.WithEncoder(encoder))))
	if err != nil {
		fmt.Println("new config failed")
		return
	}
	err = conf.Load(file.NewSource(
		file.WithPath("./resource/bootstrap.yaml"),
		source.WithEncoder(encoder)))
	if err != nil {
		fmt.Println("file load failed")
		return
	}
	// define our own host type
	type SipConfig struct {
		Serial string `json:"serial"`
		Realm  string `json:"realm"`
	}

	var sipconfig SipConfig

	fmt.Printf("values=%v \n", conf.Map())
	bs := conf.Get("sip").Bytes()
	fmt.Println(string(bs))
	err = conf.Get("sip").Scan(&sipconfig)
	if err != nil {
		fmt.Println("occur error")
	}
	fmt.Println(sipconfig)
	// viper merge
	v := viper.New()
	v.SetConfigType("json")
	err = v.MergeConfig(bytes.NewBuffer(conf.Bytes()))
	if err != nil {
		fmt.Println("viper merge config failed")
		return
	}
	serial := v.GetString("sip.serial")
	fmt.Printf("serial=%s \n", serial)
	_ = v.UnmarshalKey("sip", &sipconfig)
	fmt.Println(sipconfig)
	boot.ListenSignal()
}
