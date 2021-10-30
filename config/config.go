package config

import (
	"github.com/pkg/errors"
	"go-xstep/pkg/cache/xredis"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Env string

type Config struct {
	App    Application   `yaml:"app"`
	Port   Port          `yaml:"port"`
	Redis  xredis.Config `yaml:"redis"`
	Logger Logger        `yaml:"logger"`
}

type Application struct {
	Name string `yaml:"name"`
}

type Port struct {
	HTTPAddr  string `yaml:"tcpAddr"`
	AdminAddr string `yaml:"adminAddr"`
}

type Logger struct {
	File string `yaml:"file"`
}

func New(env Env) (c *Config) {
	c = &Config{}
	content, err := ioutil.ReadFile(string(env))
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(content, c); err != nil {
		//打印堆栈信息
		panic(errors.WithStack(err))
	}
	return
}
