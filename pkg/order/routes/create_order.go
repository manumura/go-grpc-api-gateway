package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manumura/go-grpc-api-gateway/pkg/order/pb"
)

type CreateOrderRequestBody struct {
	ProductId int64 `json:"productId"`
	Quantity  int64 `json:"quantity"`
}

func CreateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	b := CreateOrderRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, _ := ctx.Get("userId")
	fmt.Printf("type of userId is: %T\n", userId)

	u, ok := userId.(int64)
	if !ok {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	fmt.Printf("type of u is: %T\n", u)

	res, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		ProductId: b.ProductId,
		Quantity:  b.Quantity,
		UserId:    u,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
