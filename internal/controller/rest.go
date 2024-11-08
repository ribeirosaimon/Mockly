package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/Mockly/internal/service"
)

type RestController struct {
	service service.RestRoute
}

func NewRestController() *RestController {
	return &RestController{}

}

func (r *RestController) RestDinamic(c *gin.Context) {
	ctx := context.Background()
	_, err := r.service.AllRouters(ctx)
	if err != nil {

	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Endpoint: %s", "RestDinamic")})
}
