package main

import (
	"fmt"
	"golangpractice/section-4/todo-api/controller"
	"golangpractice/section-4/todo-api/model"

	_ "github.com/go-sql-driver/mysql" //mysql driver. the _ is for importing a package solely for its side-effects.

	"log"
	"net/http"
)

func main() {
	mux := controller.Register() //create new mux
	db := model.Connect()        // connect it to database
	defer db.Close()             //delayed and will run after the scope is done
	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", mux)) //listen to localhost port 3000 and pass mux to it
}
