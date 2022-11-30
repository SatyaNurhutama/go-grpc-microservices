package services

import (
	"context"
	"go-grpc-book-svc/pkg/db"
	"go-grpc-book-svc/pkg/models"
	"go-grpc-book-svc/pkg/pb"
	"net/http"
)

type Server struct {
	H db.Handler
	pb.UnsafeBookServiceServer
}

func (s *Server) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
	var book models.Book

	book.Title = req.Title
	book.Description = req.Description
	book.Author = req.Author

	if result := s.H.DB.Create(&book); result.Error != nil {
		return &pb.CreateBookResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.CreateBookResponse{
		Status: http.StatusOK,
		Id:     book.Id,
	}, nil
}

func (s *Server) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	var book models.Book

	if result := s.H.DB.First(&book, req.Id); result.Error != nil {
		return &pb.FindOneResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	data := &pb.FindOneData{
		Id:          book.Id,
		Title:       book.Title,
		Description: book.Description,
		Author:      book.Author,
	}

	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

func (s *Server) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	var book models.Book

	if result := s.H.DB.Delete(&book, req.Id); result.Error != nil {
		return &pb.DeleteBookResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.DeleteBookResponse{
		Status: http.StatusOK,
	}, nil
}
