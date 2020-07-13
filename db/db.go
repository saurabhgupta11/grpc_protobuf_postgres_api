package db

import (
	"context"
	"database/sql"

	"../dbconnection"
)

// Db global pointer for database access in API struct methods
var dbaccess *sql.DB

func init() {
	dbaccess = dbconnection.ConnectDB()
}

// Server struct
type Server struct{}

// GetDB gets data from the database and returns to the client
func (s *Server) GetDB(ctx context.Context, message *LimitOffset) (*Rows, error) {

	var reply []*SingleRow

	sqlStatement := `SELECT * from users
		ORDER BY id
		LIMIT $1
		OFFSET $2
	`
	rows, err := dbaccess.Query(sqlStatement, message.Limit, message.Offset)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		record := &SingleRow{
			Id:        0,
			Age:       0,
			Firstname: "",
			Lastname:  "",
			Email:     "",
		}
		err = rows.Scan(&record.Id, &record.Age, &record.Firstname, &record.Lastname, &record.Email)
		if err != nil {
			// handle this error
			panic(err)
		}
		reply = append(reply, record)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return &Rows{Rows: reply}, nil
}

// Insert adds a row into the database
func (s *Server) Insert(ctx context.Context, message *SingleRow) (*SingleRow, error) {
	sqlStatement := `
		INSERT INTO
		users (age, first_name, last_name, email)
		VALUES ($1, $2, $3, $4)
		RETURNING id,age,first_name,last_name,email
	`

	// data will be feeded into the reply with scan in QueryRow
	reply := &SingleRow{
		Id:        0,
		Age:       0,
		Firstname: "",
		Lastname:  "",
		Email:     "",
	}

	err := dbaccess.QueryRow(sqlStatement, message.Age, message.Firstname, message.Lastname, message.Email).Scan(&reply.Id, &reply.Age, &reply.Firstname, &reply.Lastname, &reply.Email)
	if err != nil {
		panic(err)
	}

	return reply, nil
}
