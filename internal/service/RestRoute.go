package service

import (
	"context"

	"github.com/ribeirosaimon/Mockly/internal/entity"
)

type RestRoute interface {
	AllRouters(ctx context.Context) ([]entity.Router, error)
}
type restRoute struct {
}

func NewRestRoute() *restRoute {
	return &restRoute{}
}

func (r *restRoute) AllRouters(ctx context.Context) ([]entity.Router, error) {
	return []entity.Router{}, nil
}
