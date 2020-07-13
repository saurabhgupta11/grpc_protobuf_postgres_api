package main

import (
	"context"
	"fmt"
	"log"

	"../db"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error")
	}

	defer conn.Close()

	c := db.NewDatabaseServiceClient(conn)

	message := db.LimitOffset{
		Limit:  10,
		Offset: 0,
	}

	response, err := c.GetDB(context.Background(), &message)

	for i, s := range response.Rows {
		fmt.Println(i, s)
	}
}
