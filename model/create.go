package model

import (
	"fmt"
)

func CreateTodo(name, todo string) error {
	//localhost:3000/ . Method = POST . JSON in Body
	insertQ, err := con.Query("INSERT INTO TODO VALUES (?, ?)", name, todo) //run a query to insert from JSON in http post body
	//con refering to model>connect line 20
	//assign info about result of the query and error to var insertQ and err
	defer insertQ.Close() //close when scope finished
	if err != nil {       //if error is occured
		fmt.Println(err)
		return err
	}
	return nil //will return nil bcs its a procedure not function (does not return value),
	//but print the error instead if it occured (line 14)
}

func DeleteTodo(name string) error {
	//localhost:3000/CodeChef-Vit. METHOD = DELETE. Params from URL
	insertQ, err := con.Query("DELETE FROM TODO WHERE name=?", name) //run a query DELETE where name is = name (from URL)
	//con referring to model>connect line 20
	//assign info about result of the query and error to var insertQ and err
	defer insertQ.Close() //close when scope finished
	if err != nil {       //if error is occured
		fmt.Println(err)
		return err
	}
	return nil //will return nil bcs its a procedure not function (does not return value),
	//but print the error instead if it occured (line 14)
}
