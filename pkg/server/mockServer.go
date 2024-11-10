package server

type MockEnvironment struct {
	MongoHost     string
	MongoDatabase string

	RedisHost     string
	RedisDatabase string
}

func NewMockEnvironment(mock MockEnvironment) {
	env.Mongo = DbConfig{
		Host:     mock.MongoHost,
		Database: mock.MongoDatabase,
	}
	env.Redis = DbConfig{
		Host:     mock.RedisHost,
		Database: mock.RedisDatabase,
	}
}
