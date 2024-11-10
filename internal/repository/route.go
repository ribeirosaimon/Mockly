package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/ribeirosaimon/Mockly/internal/entity"
	"github.com/ribeirosaimon/Mockly/pkg/database"
	"github.com/ribeirosaimon/Mockly/pkg/tlog"
	"go.mongodb.org/mongo-driver/bson"
)

const routesCollection = "routes"

type Route interface {
	GetByID(ctx context.Context, id string) (*entity.Router, error)
	CreateRoute(ctx context.Context, id *entity.Router) error
}

type route struct {
	collection string
	mongo      database.Mongo
}

func NewRouteRepository() *route {
	return &route{
		collection: routesCollection,
		mongo:      database.GetMongoConnection(),
	}
}

func (r *route) GetByID(ctx context.Context, id string) (*entity.Router, error) {

	var router routeDao
	result := r.mongo.GetConnection().Collection(r.collection).FindOne(ctx,
		bson.M{"_id": id, "deleted": false},
	)

	if err := result.Decode(&router); err != nil {
		return nil, err
	}

	return router.toEntity(), nil
}

func (r *route) CreateRoute(ctx context.Context, route *entity.Router) error {
	routeToDb := routeDao{
		Id:        route.GetId(),
		Name:      route.GetName(),
		Path:      route.GetPath(),
		CreatedAt: time.Now(),
	}

	inserted, err := r.mongo.GetConnection().Collection(r.collection).InsertOne(ctx, routeToDb)
	if err != nil {
		tlog.Error("RouteRepository.CreateRoute", err.Error())
		return err
	}
	tlog.Info("RouteRepository.CreateRoute", fmt.Sprintf("inserted in mongo id: %d", inserted))
	if inserted.InsertedID != nil {
		return nil
	}
	err = errors.New("error while inserting new route")
	tlog.Error("RouteRepository.CreateRoute", err.Error())
	return err
}

func (router *routeDao) toEntity() *entity.Router {
	return entity.TransformToRoute(router.Id, router.Path, router.Name, router.CreatedAt, router.UpdatedAt)
}

type routeDao struct {
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
	Id        string    `bson:"_id"`
	Path      string    `bson:"path"`
	Name      string    `bson:"name"`
	Deleted   bool      `bson:"deleted"`
}
