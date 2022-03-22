package controller

//first endpoint called "ping"

import (
	"encoding/json"
	"golangpractice/section-4/todo-api/views"
	"net/http"
)

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { //w object for response write, r to handle request
		if r.Method == http.MethodGet { //if method = get
			data := views.Response{ //assign to variable data, a struct that has code and body
				Code: http.StatusOK,
				Body: "pong",
			}
			w.WriteHeader(http.StatusOK)    //w (object to response write) , write a header status OK
			json.NewEncoder(w).Encode(data) //encode the data(line 14) to JSON with w as object
		}
	}
}
