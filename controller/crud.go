package controller

import (
	"encoding/json"
	"fmt"
	"golangpractice/section-4/todo-api/model"
	"golangpractice/section-4/todo-api/views"
	"net/http"
)

func crud() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { //w object for response write, r to handle request
		if r.Method == http.MethodPost { //if http method = post
			//localhost:3000/ . Method = POST . JSON in Body
			//take data and save it
			data := views.PostRequest{}           //assign to variable data, a struct called PostRequest that has name and todo
			json.NewDecoder(r.Body).Decode(&data) //decode the body (from http JSON) to data (struct)
			fmt.Println(data)
			if err := model.CreateTodo(data.Name, data.Todo); err != nil { //if error found
				w.Write([]byte("Some error"))
				return
			}
			w.WriteHeader(http.StatusCreated) //write http status created as header
			json.NewEncoder(w).Encode(data)   //encode with w. from data (struct) to json
		} else if r.Method == http.MethodGet { //if http method = get
			//localhost:3000/?name=CodeChef-Vit . Method = GET . Params from URL
			name := r.URL.Query().Get("name") // assign name from url params "name" to variable name
			if name != "" {                   //has parameters
				data, err := model.ReadByName(name) //refering to model>create.go
				if err != nil {
					w.Write([]byte(err.Error())) //if error found, write error
				}
				w.WriteHeader(http.StatusOK)    //write http status ok as header
				json.NewEncoder(w).Encode(data) //encode with w. from data (struct) to json
			} else { //no parameters
				data, err := model.ReadAll()
				if err != nil {
					w.Write([]byte(err.Error())) //if error found, write error
				}
				w.WriteHeader(http.StatusOK)    //write http status ok as header
				json.NewEncoder(w).Encode(data) //encode with w. from data (struct) to json
			}
		} else if r.Method == http.MethodDelete { //if http method = delete
			//localhost:3000/CodeChef-Vit . Method = DELETE . Params from URL
			name := r.URL.Path[1:]                         //take url params starting from / to end in this case : CodeChef-Vit
			if err := model.DeleteTodo(name); err != nil { // if error occured
				w.Write([]byte("Some error"))
				return
			}
			w.WriteHeader(http.StatusOK)       //write http status ok as a header
			json.NewEncoder(w).Encode(struct { //encode a struct to json as object w. the status is Item Deleted
				Status string `json:status`
			}{"Item deleted"})
		}
	}
}
