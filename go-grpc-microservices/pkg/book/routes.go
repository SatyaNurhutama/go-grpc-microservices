package book

import (
	"go-grpc-microservices/pkg/auth"
	"go-grpc-microservices/pkg/book/routes"
	"go-grpc-microservices/pkg/config"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/book")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateBook)
	routes.GET("/:id", svc.FindOne)
	routes.DELETE("/:id", svc.DeleteBook)

}

func (svc *ServiceClient) CreateBook(ctx *gin.Context) {
	routes.CreateBook(ctx, svc.Client)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}

func (svc *ServiceClient) DeleteBook(ctx *gin.Context) {
	routes.DeleteBook(ctx, svc.Client)
}
