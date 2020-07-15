package main

import (
	"crypto/tls"
	"log"
	"net"

	"./db"
	_ "github.com/lib/pq" // features specific to postgresql
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair("cert/server-cert.pem", "cert/server-key.pem")
	if err != nil {
		return nil, err
	}

	// create the credentials
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Fatal to listen on port 9000: %v", err)
	}

	s := db.Server{}

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("FAiled to load TLS credentials", err)
	}

	grpcServer := grpc.NewServer(
		grpc.Creds(tlsCredentials),
	)

	db.RegisterDatabaseServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed %v", err)
	}
}
