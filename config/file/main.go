package main

import (
	"bytes"
	"fmt"
	"github.com/asim/go-micro/plugins/config/encoder/yaml/v4"
	"github.com/spf13/viper"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/reader/json"
	"go-micro.dev/v4/config/source"
	"go-micro.dev/v4/config/source/file"
)

func main() {
	readFromYaml()
}
func readFromYaml() {
	// 定义yaml解析器
	enc := yaml.NewEncoder()
	r := json.NewReader(reader.WithEncoder(enc))
	conf, err := config.NewConfig(config.WithReader(r))
	if err != nil {
		return
	}

	fileSource := file.NewSource(file.WithPath("./resource/bootstrap.yaml"), source.WithEncoder(enc))
	err = conf.Load(fileSource)
	if err != nil {
		panic(err)
	}
	fmt.Println("data", conf.Map())

	type sip struct {
		Serial string
		Realm  string
	}
	s := sip{}

	jsonValue := conf.Get("cc", "sip")
	_ = jsonValue.Scan(&s)
	fmt.Printf("scan=%v \n", s)

	v := viper.New()
	v.SetConfigType("json")
	_ = v.MergeConfig(bytes.NewBuffer(jsonValue.Bytes()))
	_ = v.Unmarshal(&s)

	fmt.Printf("unmarshal=%v \n", s)

	serial := v.GetInt("serial")
	fmt.Println(serial)

	// config 使用micro ,具备watch机制
	//需要序列化时，如果存在数字字符串的情况使用viper进行替换
}
