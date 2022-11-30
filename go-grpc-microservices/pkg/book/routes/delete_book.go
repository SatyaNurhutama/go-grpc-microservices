package routes

import (
	"context"
	"go-grpc-microservices/pkg/book/pb"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteBook(ctx *gin.Context, c pb.BookServiceClient) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	res, err := c.DeleteBook(context.Background(), &pb.DeleteBookRequest{
		Id: id,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
