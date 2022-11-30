package routes

import (
	"context"
	"go-grpc-microservices/pkg/book/dto"
	"go-grpc-microservices/pkg/book/pb"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook(ctx *gin.Context, c pb.BookServiceClient) {
	body := dto.CreateBookRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateBook(context.Background(), &pb.CreateBookRequest{
		Title:       body.Title,
		Description: body.Description,
		Author:      body.Author,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
