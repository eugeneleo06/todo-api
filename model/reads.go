package model

import "golangpractice/section-4/todo-api/views"

func ReadAll() ([]views.PostRequest, error) { //read as array from struct PostRequest
	rows, err := con.Query("SELECT * FROM TODO") //run a query to select all from todo
	//assign info about result of the query and error to var rows and err
	if err != nil { //if error occured
		return nil, err
	}
	todos := []views.PostRequest{} //assign struct postrequest from views to array variable todos
	for rows.Next() {              //for each rows, from data variable, the Name and Todo elements and append
		//each row (var data) to array todos
		data := views.PostRequest{}       //assign postrequest struct to var data
		rows.Scan(&data.Name, &data.Todo) //scan for key Name and Todo
		todos = append(todos, data)       //append data (each rows) to array variable todos (declared in line 11)
	}
	return todos, nil //return the array from ReadAll it will return data whole table
}

func ReadByName(name string) ([]views.PostRequest, error) { //read as array from struct PostRequest
	//localhost:3000/?name=CodeChef-Vit . Method = GET . Params from URL
	rows, err := con.Query("SELECT * FROM TODO WHERE name = ?", name) //run a query to select which name from Parameters GET in URL
	//assign the info about rows and err into it's variable
	if err != nil { //if error occured
		return nil, err
	}
	todos := []views.PostRequest{} //assign struct postrequest from views to array variable todos
	for rows.Next() {              //for each rows, from data variable, the Name and Todo elements and append
		//each row (var data) to array todos
		data := views.PostRequest{}       //assign postrequest struct to var data
		rows.Scan(&data.Name, &data.Todo) //scan for key Name and Todo
		todos = append(todos, data)       //append data (each rows) to array variable todos (declared in line 28)
	}
	return todos, nil //return the array from ReadByName it will return data which name is the same as the Params GET from URL
}
