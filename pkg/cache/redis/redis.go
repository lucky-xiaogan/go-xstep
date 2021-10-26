package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
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

func Dail(addr, password string, options ...func(*Config)) (*Config, error) {
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

	return &c, nil
}

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

var (
	rdb *redis.Client
)

func NewRedis(c *Config) error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:16379",
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	return err
}
