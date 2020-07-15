package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"../db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/proto"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// load certificate of CA who signed web server's certificate
	pemServerCA, err := ioutil.ReadFile("../cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	// create the credentials
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server's ca certificate")
	}

	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	var conn *grpc.ClientConn

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("FAiled to load TLS credentials", err)
	}

	conn, err = grpc.Dial(":9000", grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatal("Error")
	}

	defer conn.Close()

	c := db.NewDatabaseServiceClient(conn)

	message := db.LimitOffset{
		Limit:  30000,
		Offset: 0,
	}

	// measuring time to retrieve data from the server
	start := time.Now()

	response, err := c.GetDB(context.Background(), &message)

	out, e := proto.Marshal(response)
	if e != nil {
		panic(e)
	}

	// fmt.Println(string(out))
	out, e = ioutil.ReadFile("data")
	if e != nil {
		log.Fatal("failed", e)
	}

	newMessage := &db.Rows{}
	e = proto.Unmarshal(out, newMessage)
	if e != nil {
		panic(e)
	}

	fmt.Println(newMessage)

	elapsed := time.Since(start)

	// for i, s := range response.Rows {
	// 	fmt.Println(i, s)
	// }

	fmt.Println("time taken : ", elapsed)

	// trying to insert 1000 rows in the database
	// for i := 1000; i < 30000; i++ {
	// 	record := db.SingleRow{
	// 		Id:        0,
	// 		Age:       (int64)(i),
	// 		Firstname: "a",
	// 		Lastname:  "b",
	// 		Email:     strconv.Itoa(i) + "@gmail.com",
	// 	}

	// 	_, err := c.Insert(context.Background(), &record)
	// 	if err != nil {
	// 		log.Fatal("error occured", err)
	// 	}
	// }
}
