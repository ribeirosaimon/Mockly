package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/Mockly/pkg/server"
)

func main() {
	if myEnv := os.Getenv("ENVIRONMENT"); myEnv != "" {
		server.StartEnv(server.Environment(myEnv))
	}

	configs := server.GetEnvironment()
	if configs.Env == "" {
		log.Fatal("Environment variable not set")
	}

	router := gin.Default()

	server.InitControllers(router)

	router.Run(":8080")
}
