package conf

import (
	"os"

	"github.com/larspensjo/config"
)

var Conf, UrlConf *config.Config

func BuildConfig() {
	conf, err := config.ReadDefault("config/app.conf") //将配置文件加载到Conf对象中
	if err != nil {
		os.Exit(1)
	}
	Conf = conf
}

func init() {
	BuildConfig()
	return
}
