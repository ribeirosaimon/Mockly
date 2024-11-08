package database

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	onceMongo       sync.Once
	mongoConn       mongoConnection
	mongoDefaultUrl = "mongodb://localhost:27017"
)

type Mongo interface {
	GetConnection() *mongo.Database
}

// Option was a function optional pattern
type MongoOption func(*mongoConnection)

func WithUrl(url string) MongoOption {
	return func(a *mongoConnection) {
		a.url = url
	}
}

func WithDatabase(db string) MongoOption {
	return func(a *mongoConnection) {
		a.database = db
	}
}

func NewConnMongo(opts ...MongoOption) *mongoConnection {
	ctx := context.Background()
	mongoConn = mongoConnection{
		url: mongoDefaultUrl,
	}
	mongoConn.conn(ctx)

	for _, opt := range opts {
		opt(&mongoConn)
	}
	if mongoConn.database == "" {
		panic("Need to set DATABASE")
	}
	onceMongo.Do(func() {
		mongoConn.mongo = mongoConn.conn(ctx)
	})
	return &mongoConn
}

type mongoConnection struct {
	url      string
	database string
	mongo    *mongo.Database
}

func (c *mongoConnection) GetConnection() *mongo.Database {
	return c.mongo
}

func (c *mongoConnection) conn(ctx context.Context) *mongo.Database {
	clientOptions := options.Client().ApplyURI(c.url)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(c.database)
}
