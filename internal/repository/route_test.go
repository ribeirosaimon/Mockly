package repository

import (
	"context"
	"testing"

	"github.com/ribeirosaimon/Mockly/internal/entity"
	"github.com/ribeirosaimon/Mockly/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestNewRouteRepository(t *testing.T) {
	database.NewMongo(database.WithUrl("mongodb://localhost:27017"), database.WithDatabase("mockly_test"))

	repository := NewRouteRepository()
	ctx := context.Background()

	for _, singleTest := range []struct {
		testName string
	}{
		{
			testName: "Need save my route",
		},
	} {
		t.Run(singleTest.testName, func(t *testing.T) {
			router := NewRouter()
			err := repository.CreateRoute(ctx, router)
			assert.NotNil(t, err)

			routerDb, err := repository.GetByID(ctx, router.GetId())
			assert.Nil(t, err)
			assert.Equal(t, router.GetId(), routerDb.GetId())
			assert.Equal(t, router.GetName(), routerDb.GetName())
			assert.Equal(t, router.GetPath(), routerDb.GetPath())
		})
	}
}

func NewRouter() *entity.Router {
	return entity.NewRoute("test", "testName")
}
