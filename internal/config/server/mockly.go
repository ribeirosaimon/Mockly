package server

import "github.com/gin-gonic/gin"

type Mockly struct {
}

func NewMockly() *Mockly {
	return &Mockly{}
}

func (m *Mockly) StartServer() {
	router := gin.Default()

	InitControllers(router)

	router.Run(":8080")
}
