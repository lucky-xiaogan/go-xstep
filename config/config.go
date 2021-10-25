package config

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Env string

type Config struct {
	App   Application `yaml:"app"`
	Port  Port        `yaml:"port"`
	Redis RedisConfig `yaml:"redis"`
}

type Application struct {
	Name string `yaml:"name"`
}

type Port struct {
	HTTPAddr  string `yaml:"tcpAddr"`
	AdminAddr string `yaml:"adminAddr"`
}

type RedisConfig struct {
	Addr         string `yaml:"host"`
	Password     string `yaml:"password"`
	DB           int    `yaml:"db"`
	MaxRetries   int    `yaml:"maxRetries"`
	PoolSize     int    `yaml:"poolSize"`
	MinIdleConns int    `yaml:"minIdleConns"`
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
