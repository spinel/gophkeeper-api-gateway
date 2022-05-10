package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spinel/gophkeeper-api-gateway/pkg/auth"
	"github.com/spinel/gophkeeper-api-gateway/pkg/config"
	"github.com/spinel/gophkeeper-api-gateway/pkg/storage"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	storage.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
