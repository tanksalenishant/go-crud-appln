package main

import (
	"context"
	"fmt"
	"log"

	// "rest1/internal/domain"
	// "rest1/internal/repository"
	"github.com/jackc/pgx/v5"
)

func main() {
	// Replace these values with your PostgreSQL details
	username := "nishant"
	password := "nishant"
	host := "db"
	port := "5432"
	database := "nishant"
	fmt.Println("Connecting to .....")
	// Connection string
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", username, password, host, port, database)

	// Establish a connection to the PostgreSQL database
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())
	// Check the connection
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to PostgreSQL!")

	// Perform database operations here...

	// Example: Querying the version of the PostgreSQL server
	// var version string
	// err = conn.QueryRow(context.Background(), "SELECT version()").Scan(&version)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//  _ , err = conn.Exec(context.Background() , "CREATE TABLE users (id INT PRIMARY KEY,name VARCHAR ( 50 )  NOT NULL,password VARCHAR ( 50 ) NOT NULL,accountNo INT);")
	//  if err != nil{
	// 	fmt.Println(err)
	//  }
	// var user domain.User;
	// user.ID = 2003
	// user.AccountNo = 123
	// user.Name = "NMMMM"
	// user.Password = "abcd"
	// user , err = repository.NewUserRepo(conn).Create(user)
	// fmt.Println("PostgreSQL version:", version)
	// users, err := repository.NewUserRepo(conn).GetAll(conn)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	// Print the users to the response.
	// 	for _, user := range users {
	// 		fmt.Printf("ID: %d, Name: %s, AccountNo: %d, Password: %s\n", user.ID, user.Name, user.AccountNo, user.Password)
	// 	}

	// //8000 port in router

}