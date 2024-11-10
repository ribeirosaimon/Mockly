package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Mockly struct {
	Port string
}

func NewMockly(options ...MocklyOpt) *Mockly {
	m := &Mockly{
		Port: "8080",
	}
	for _, opt := range options {
		opt(m)
	}
	return m
}

// MocklyOpt was a function optional pattern
type MocklyOpt func(*Mockly)

func WithPort(port string) MocklyOpt {
	return func(m *Mockly) {
		m.Port = port
	}
}

func (m *Mockly) StartServer() {
	router := gin.Default()

	router.Run(fmt.Sprintf(":%s", m.Port))
}
