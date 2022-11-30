package book

import (
	"go-grpc-microservices/pkg/book/pb"
	"go-grpc-microservices/pkg/config"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.BookServiceClient
}

func InitServiceClient(c *config.Config) pb.BookServiceClient {
	conn, err := grpc.Dial(c.BookSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to: %s", err)
	}

	return pb.NewBookServiceClient(conn)
}
