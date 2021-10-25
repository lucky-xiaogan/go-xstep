package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//解析yml文件
type BaseInfo struct {
	Port   string      `yaml:"port"`
	Ip     string      `yaml:"ip"`
	Host   string      `yaml:"host"`
	Spring RedisEntity `yaml:"spring"`
}

type RedisEntity struct {
	Redis RedisData `yaml:"redis"`
}

type RedisData struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DataBase string `yaml:"dataBase"`
	Timeout  string `yaml:"timeout"`
}

func (c *BaseInfo) GetConf() *BaseInfo {
	yamlFile, err := ioutil.ReadFile(".config.yml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
