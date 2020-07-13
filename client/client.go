package main

import (
	"context"
	"fmt"
	"log"
	"time"

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
		Limit:  3,
		Offset: 0,
	}

	// measuring time to retrieve data from the server
	start := time.Now()

	response, err := c.GetDB(context.Background(), &message)

	elapsed := time.Since(start)

	for i, s := range response.Rows {
		fmt.Println(i, s)
	}

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
