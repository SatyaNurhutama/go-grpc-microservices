package main

import (
	"fmt"
	"go-grpc-book-svc/pkg/config"
	"go-grpc-book-svc/pkg/db"
	"go-grpc-book-svc/pkg/pb"
	"go-grpc-book-svc/pkg/services"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Book svc on:", c.Port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	pb.RegisterBookServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
