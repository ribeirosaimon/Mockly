package cache

import (
	"context"

	"github.com/ribeirosaimon/Mockly/pkg/database"
)

type Route interface {
	GetRoute(context.Context, string) string
}

type route struct {
	redis database.Redis
}

func NewRouteCache() *route {
	return &route{
		redis: database.NewRedisConnection(),
	}
}

func (r *route) GetRoute(ctx context.Context, key string) string {
	value := r.redis.GetConnection().Get(ctx, key)
}
