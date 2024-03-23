package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

// serverImpl реализует ваш gRPC-сервер
type serverImpl struct{}

// CheckHash реализует RPC-метод CheckHash
func (s *serverImpl) CheckHash(ctx context.Context, req *proto.CheckHashRequest) (*proto.CheckHashResponse, error) {
	// Реализуйте вашу логику здесь
	return &proto.CheckHashResponse{}, nil
}

// GetHash реализует RPC-метод GetHash
func (s *serverImpl) GetHash(ctx context.Context, req *proto.GetHashRequest) (*proto.GetHashResponse, error) {
	// Реализуйте вашу логику здесь
	return &proto.GetHashResponse{}, nil
}

// CreateHash реализует RPC-метод CreateHash
func (s *serverImpl) CreateHash(ctx context.Context, req *proto.CreateHashRequest) (*proto.CreateHashResponse, error) {
	// Реализуйте вашу логику здесь
	return &proto.CreateHashResponse{}, nil
}

func main() {
	// Создаем новый gRPC сервер
	grpcServer := grpc.NewServer()

	// Регистрируем ваш сервис на сервере
	proto.RegisterHashingServiceServer(grpcServer, &serverImpl{})

	// Указываем порт, на котором будет слушать сервер
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Запускаем сервер
	fmt.Println("Server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
