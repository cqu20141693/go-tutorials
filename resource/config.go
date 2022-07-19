package resource

import (
	"github.com/cqu20141693/go-tutorials/utils"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var App Config

// Init 读取全局配置文件
func init() {
	paths := make([]string, 0, 30)
	paths = append(paths, "./resource/default.yaml")
	_, err := os.Stat("/etc/mongo.yaml")
	if err == nil || os.IsExist(err) {
		paths = append(paths, "/etc/topology.yaml")
	} else {
		yamls, err := WalkDir("./configs", ".yaml")
		if err == nil {
			paths = append(paths, yamls...)
		}
	}
	for _, c := range paths {
		data, err := ioutil.ReadFile(c)
		if err == nil {
			yaml.Unmarshal(data, &App)
			log.Debug().Caller().Str("func", "config.Init").Msgf("Read config: config=%v, app=%v", c, App)
		} else {
			log.Warn().Err(err).Msgf("Read config constant.")
		}
	}

	getEnvConfig()
	if App.Port == 0 {
		App.Port = 8200
	}

	log.Info().Msgf("App config: %v", App)
}

func getEnvConfig() {

	text := os.Getenv("MONGO_DATABASE")
	if text != "" {
		App.Mongo.Database = text
	}

	text = os.Getenv("MONGO_USER")
	if text != "" {
		App.Mongo.User = text
	}

	text = os.Getenv("MONGO_PASSWORD")
	if text != "" {
		App.Mongo.Password = text
	}

	text = os.Getenv("MONGO_MAXCONNECTTIONS")
	if text != "" {
		App.Mongo.MaxConnections = utils.Int(text)
	}

	text = os.Getenv("MONGO_TIMEOUT")
	if text != "" {
		App.Mongo.Timeout = utils.Int(text)
	}

	text = os.Getenv("MONGO_MECHANISM")
	if text != "" {
		App.Mongo.Mechanism = text
	}
	text = os.Getenv("MONGO_AUTHSOURCE")
	if text != "" {
		App.Mongo.AuthSource = text
	}
	text = os.Getenv("MONGO_DEBUG")
	if text == "true" {
		App.Mongo.Debug = true
	}
}

// WalkDir 获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func WalkDir(dirPth, suffix string) ([]string, error) {
	files := make([]string, 0, 30)
	_, err := os.Stat(dirPth)
	if err == nil || os.IsExist(err) {
		suffix = strings.ToUpper(suffix)
		err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error {
			// 忽略目录
			if fi.IsDir() {
				return nil
			}

			if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
				files = append(files, filename)
			}
			return nil
		})
	}

	return files, err
}
