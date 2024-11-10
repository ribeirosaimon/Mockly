package database

import (
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	oncePgsql       sync.Once
	redisConn       redisConnection
	pgsqlDefaultUrl = "localhost:6379"
)

type RedisOption func(*redisConnection)

func RedisWithUrl(url string) RedisOption {
	return func(a *redisConnection) {
		a.url = url
	}
}

type Redis interface {
	GetConnection() *redis.Client
}

type redisConnection struct {
	url      string
	password string
	database int
	redis    *redis.Client
}

func NewRedisConnection(opts ...RedisOption) *redisConnection {
	redisConn = redisConnection{
		url: pgsqlDefaultUrl,
	}
	for _, opt := range opts {
		opt(&redisConn)
	}

	oncePgsql.Do(func() {
		conn := redisConn.conn()
		redisConn.redis = conn
	})

	return &redisConn
}

func (c *redisConnection) conn() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.url,
		Password: c.password,
		DB:       c.database,
	})
}

func (c *redisConnection) GetConnection() *redis.Client {
	return c.redis
}
