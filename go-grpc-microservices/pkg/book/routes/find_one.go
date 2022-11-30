package routes

import (
	"context"
	"go-grpc-microservices/pkg/book/pb"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindOne(ctx *gin.Context, c pb.BookServiceClient) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	res, err := c.FindOne(context.Background(), &pb.FindOneRequest{
		Id: int64(id),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
