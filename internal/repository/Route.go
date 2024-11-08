package repository

import (
	"github.com/ribeirosaimon/Mockly/internal/config/database"
)

type Route interface {
}

type route struct {
	mongo database.Mongo
}

func NewRouteRepository() *route {
	return &route{
		mongo: database.NewConnMongo(),
	}
}
