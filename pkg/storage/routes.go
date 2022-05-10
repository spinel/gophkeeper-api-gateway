package storage

import (
	"github.com/gin-gonic/gin"
	"github.com/spinel/gophkeeper-api-gateway/pkg/auth"
	"github.com/spinel/gophkeeper-api-gateway/pkg/config"
	"github.com/spinel/gophkeeper-api-gateway/pkg/storage/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("entity")
	routes.Use(a.AuthRequired)
	routes.POST("/create", svc.CreateEntity)
}

func (svc *ServiceClient) CreateEntity(ctx *gin.Context) {
	routes.CreateEntity(ctx, svc.Client)
}
