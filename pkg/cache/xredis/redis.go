package xredis

import (
	"github.com/go-redis/redis/v8"
)

type Config struct {
	Addr         string `yaml:"host"`
	Password     string `yaml:"password"`
	DB           int    `yaml:"db"`
	MaxRetries   int    `yaml:"maxRetries"`
	PoolSize     int    `yaml:"poolSize"`
	MinIdleConns int    `yaml:"minIdleConns"`
}

type Option func(*Config)


func DialDatabase(db int) Option {
	return func(config *Config) {
		config.DB = db
	}
}

func DailPassword(password string) Option {
	return func(config *Config) {
		config.Password = password
	}
}

func DailMaxRetries(maxRetries int) Option {
	return func(config *Config) {
		config.MaxRetries = maxRetries
	}
}

func DailPoolSize(poolSize int) Option {
	return func(config *Config) {
		config.PoolSize = poolSize
	}
}

func DailMinIdleConns(minIdleConns int) Option {
	return func(config *Config) {
		config.MinIdleConns = minIdleConns
	}
}

func NewRedis(addr, password string, options ...func(*Config)) *redis.Client {
	c := Config{
		Addr:         addr,
		Password:     password,
		DB:           0,
		MaxRetries:   3,
		PoolSize:     10,
		MinIdleConns: 10,
	}

	for _, option := range options {
		option(&c)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:       c.Addr,
		Password:   c.Password, // no password set
		DB:         c.DB,       // use default DB
		MaxRetries: c.MaxRetries,
		PoolSize:   c.PoolSize, // 连接池大小
	})

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//
	//_, err := RedisDB.Ping(ctx).Result()
	return rdb
}


