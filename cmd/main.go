package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/Egor-Golang-TSM-Course/final-project-vechnonetot/hashing"
)

func main() {

	// Создаем gRPC сервер
	server := grpc.NewServer()

	// Регистрируем сервис на сервере
	hashing.RegisterHashingServiceServer(server, service)

	// Начинаем слушать запросы на порту 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Сервис хеширования запущен")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
