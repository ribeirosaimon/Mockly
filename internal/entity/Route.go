package entity

import (
	"time"

	"github.com/ribeirosaimon/Mockly/internal/vo"
)

type Router struct {
	createdAt  time.Time
	updatedAt  time.Time
	id         vo.Id
	responseId vo.Id
	path       vo.Path
	name       vo.Name
	deleted    bool
}

func NewRoute(path, name string) *Router {
	return &Router{
		id:        vo.NewId(),
		path:      vo.NewPath(path),
		name:      vo.NewName(name),
		createdAt: time.Now(),
		deleted:   false,
	}
}

func TransformToRoute(id, path, name string, createdAt, updatedAt time.Time) *Router {
	return &Router{
		id:        vo.TransformStringInId(id),
		path:      vo.NewPath(path),
		name:      vo.NewName(name),
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (r *Router) SetResponseId(id vo.Id) {
	r.responseId = id
	r.updatedAt = time.Now()
}

func (r *Router) GetId() string {
	return r.id.GetValue()
}
func (r *Router) GetPath() string {
	return r.path.GetValue()
}

func (r *Router) GetResponseId() string {
	return r.responseId.GetValue()
}

func (r *Router) GetName() string {
	return r.name.GetValue()
}
