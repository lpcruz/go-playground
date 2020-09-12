package main

import (
	"net/http"

	"github.com/lpcruz/go-playground/projects/todo-list/api"
)

func main() {
	server := api.NewServer()
	http.ListenAndServe(":8080", server)
}
