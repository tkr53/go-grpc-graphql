package main

import (
	"log"
	"net"

	"github.com/tkr53/go-gprc-graphql/article/pb"
	"github.com/tkr53/go-gprc-graphql/article/repository"
	"github.com/tkr53/go-gprc-graphql/article/service"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	defer lis.Close()

	repository, err := repository.NewSqliteRepo()
	if err != nil {
		log.Fatalf("Failed to create splite repository: %v\n", err)
	}

	service := service.NewService(repository)

	server := grpc.NewServer()
	pb.RegisterArticleServiceServer(server, service)

	log.Println("Listening on port 50051...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
