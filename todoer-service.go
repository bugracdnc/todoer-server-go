package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type TodoService struct {
	todoRepository TodoRepository
}

func getIdFromQuery(r *http.Request) (uuid.UUID, error) {
	idQuery := chi.URLParam(r, "id")
	return uuid.Parse(idQuery)

}

func (s *TodoService) handleGet(w http.ResponseWriter, userId *uuid.UUID) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(*s.todoRepository.getAll(userId))
	if err != nil {
		fmt.Printf("error (handleGet): %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func (s *TodoService) handleGetById(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromQuery(r)
	if err != nil {
		fmt.Printf("error (getByIdHandler): %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	encodeErr := json.NewEncoder(w).Encode(*s.todoRepository.getById(id))
	if encodeErr != nil {
		fmt.Printf("error (handleGetById): %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func (s *TodoService) handleCreate(w http.ResponseWriter, r *http.Request, user *User) {
	var todo = Todo{}
	err := json.NewDecoder(r.Body).Decode(&todo)
	todo.UserId = user.Id
	if err != nil {
		fmt.Printf("error (handleCreate): %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if len(todo.Todo) < 1 {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	s.todoRepository.save(todo)
	w.WriteHeader(http.StatusCreated)
}

func (s *TodoService) handleUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromQuery(r)
	if err != nil {
		fmt.Printf("error (handleUpdate):  %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	todo := s.todoRepository.getById(id)
	if todo == nil {
		fmt.Printf("error (handleUpdate):  %s\n", "Cannot find todo")
		http.Error(w, "Cannot find todo", http.StatusNotFound)
		return
	}

	var sentTodo Todo
	decodeErr := json.NewDecoder(r.Body).Decode(&sentTodo)
	if decodeErr != nil {
		fmt.Printf("error (decodeErr): %s\n", decodeErr.Error())
		http.Error(w, decodeErr.Error(), http.StatusBadRequest)
	}
	todo.Id = id
	todo.Todo = sentTodo.Todo
	todo.Done = sentTodo.Done
	todo.UpdateDate = time.Now()

	s.todoRepository.update(*todo)
	w.WriteHeader(http.StatusOK)
}

func (s *TodoService) handleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromQuery(r)
	if err != nil {
		fmt.Println(err)
	}
	todo := s.todoRepository.getById(id)
	if todo == nil {
		fmt.Printf("error (handleUpdate):  %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	s.todoRepository.delete(id)
	w.WriteHeader(http.StatusOK)
}
