package main

import (
	"log"
	"net"

	"./db"
	_ "github.com/lib/pq" // features specific to postgresql
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Fatal to listen on port 9000: %v", err)
	}

	s := db.Server{}

	grpcServer := grpc.NewServer()

	db.RegisterDatabaseServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed %v", err)
	}
}
