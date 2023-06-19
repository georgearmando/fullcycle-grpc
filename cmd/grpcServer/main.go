package main

import (
	"database/sql"
	"net"

	"github.com/georgermando/fullcycle-gRPC/internal/database"
	"github.com/georgermando/fullcycle-gRPC/internal/pb"
	"github.com/georgermando/fullcycle-gRPC/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite") // Abre uma conexao com o DB sqlite

	if err != nil {
		panic(err)
	}

	defer db.Close() // Fecha a conexao com o db assim que a execucao do mi terminar

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb) // Cria o serviso de categoria

	grpcServer := grpc.NewServer()                                // Cria o servidor grpc
	pb.RegisterCategoryServiceServer(grpcServer, categoryService) // regista o servico ao server
	reflection.Register(grpcServer)

	// Verifica a disponibilidade da porta e atribui a lis
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	// Hablita a porta em que o server estara ouvindo
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
