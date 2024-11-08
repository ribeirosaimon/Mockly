package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	server.InitControllers(router)

	router.Run(":8080")
}
