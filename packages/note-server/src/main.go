package main

import (
	"log"
	"net"
	"note-server/pb"
	"note-server/src/db"
	"note-server/src/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":50051"

func initGrpcServer() *grpc.Server {
	grpcServer := grpc.NewServer()
	noteServer := server.NewNoteServer()

	pb.RegisterNoteServiceServer(grpcServer, noteServer)

	reflection.Register(grpcServer)

	log.Printf("🚀 gRPC server starting on port %s", port)

	return grpcServer
}

func main() {
	if err := db.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Println("Error closing database:", err)
		}
	}()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	grpcServer := initGrpcServer()

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
