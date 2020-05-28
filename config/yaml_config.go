package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

/*
yaml配置文件的读取 配置
*/
type Conf struct {
	Server struct {
		Port      string `yaml:"port"`
		EvictionS int64  `yaml:"eviction-in-s"`
	}
}

func (c *Conf) GetConf() *Conf {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err    #%v", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal:  %v", err)
	}
	return c
}
