package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB //whatever function delete, etc requires access to connection.
//con only model package has acces to the connection
//Con every package has access

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/db_todo") // _ for importing a package solely for its side-effects.
	// open database to variable db and store the error to variable "err"
	// "user:password@tcp(localhost:3306)/db_name"
	if err != nil { //if error occured
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
	con = db  //assign the db from line 14 to a variable con so it can be used in another file
	return db //return the db (referring to)
}
