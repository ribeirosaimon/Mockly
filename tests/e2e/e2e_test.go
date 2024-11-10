package e2e

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/ribeirosaimon/Mockly/pkg/server"
)

func TestE2E(t *testing.T) {
	stop := make(chan bool)

	go func() {
		server.NewMockly().StartServer()
		<-stop
		log.Println("Servidor encerrado.")
	}()

	t.Run("Have to get All routers", func(t *testing.T) {
		request, err := httpRequest(http.MethodGet, "", "/all")
		if err != nil {
			t.Fatal(err)
		}
		log.Printf("Response: %s", request)
	})

	defer func() {
		stop <- true
	}()

	// Espera um pouco para garantir que o servidor tenha tempo de encerrar
	time.Sleep(1 * time.Second)
}
