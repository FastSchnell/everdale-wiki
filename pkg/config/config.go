package config

import (
	"everdale-wiki/pkg/logger"
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var CfgInstance = new(Cfg)

func init() {
	path := flag.String("c", "conf/config.yaml", "config.yaml file path.")
	flag.Parse()
	cfg, err := getCfg(*path)
	if err != nil {
		logger.Fatal.Json(map[string]interface{}{
			"flag": "config init",
			"err":  err.Error(),
		})
	} else {
		logger.Info.Json(map[string]interface{}{
			"flag": "config init success",
			"path": *path,
			"cfg":  cfg,
		})
	}

	CfgInstance = cfg

}

func getCfg(path string) (cfg *Cfg, err error) {
	cfg = &Cfg{}
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(yamlFile, cfg)
	return
}
