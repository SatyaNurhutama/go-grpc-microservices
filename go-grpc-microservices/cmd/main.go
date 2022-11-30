package main

import (
	"go-grpc-microservices/pkg/auth"
	"go-grpc-microservices/pkg/book"
	"go-grpc-microservices/pkg/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	book.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
