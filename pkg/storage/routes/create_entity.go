package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spinel/gophkeeper-api-gateway/pkg/storage/pb"
)

type EntityType struct {
	pb.CreateEntityRequest_CreditCard
	pb.CreateEntityRequest_Account
}

type CreateEntityRequestBody struct {
	Identifier string     `json:"identifier"`
	TypeID     int64      `json:"type_id"`
	Entity     EntityType `json:"entity"`
}

func CreateEntity(ctx *gin.Context, c pb.StorageServiceClient) {
	body := CreateEntityRequestBody{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var err error
	var res *pb.CreateEntityResponse

	switch body.TypeID {
	case 1:
		res, err = c.CreateEntity(context.Background(), &pb.CreateEntityRequest{
			Identifier: body.Identifier,
			TypeID:     body.TypeID,
			Entity:     &body.Entity.CreateEntityRequest_Account,
		})
	case 2:
		res, err = c.CreateEntity(context.Background(), &pb.CreateEntityRequest{
			Identifier: body.Identifier,
			TypeID:     body.TypeID,
			Entity:     &body.Entity.CreateEntityRequest_CreditCard,
		})
	}

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
