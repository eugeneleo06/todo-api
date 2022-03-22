package controller

import (
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	//instantiate new mux
	mux.HandleFunc("/ping", ping()) //handle route /ping with func ping()
	//handle route for localhost:3000/ping
	mux.HandleFunc("/", crud()) //handle route / with func crud()
	//handle route for localhost:3000/
	return mux //return mux, referring to Main Line 15
}
