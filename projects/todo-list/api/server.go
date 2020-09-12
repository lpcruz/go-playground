package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Item struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Server struct {
	*mux.Router

	todoItems []Item
}

func NewServer() *Server {
	s := &Server{
		Router:    mux.NewRouter(),
		todoItems: []Item{},
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	fmt.Println("Starting server at http://localhost:8080")
	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	s.HandleFunc("/todo-items", s.listTodoItems()).Methods("GET")
	s.HandleFunc("/todo-items", s.createTodoItem()).Methods("POST")
	s.HandleFunc("/todo-items/{id}", s.removeTodoItem()).Methods("DELETE")
}

func (s *Server) createTodoItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i Item
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		i.ID = uuid.New()
		s.todoItems = append(s.todoItems, i)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(i); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

func (s *Server) listTodoItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.todoItems); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func (s *Server) removeTodoItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr, _ := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		for i, item := range s.todoItems {
			if item.ID == id {
				s.todoItems = append(s.todoItems[:i], s.todoItems[i+1:]...)
				break
			}
		}
	}
}
