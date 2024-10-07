package main

import (
	"fmt"
	"net/http"
)

type TodoController struct {
	todoService TodoService
}

func (t *TodoController) getHandler(w http.ResponseWriter, r *http.Request) {
	t.todoService.handleGet(w)
}

func (t *TodoController) getByIdHandler(w http.ResponseWriter, r *http.Request) {

	t.todoService.handleGetById(w, r)
}

func (t *TodoController) createHandler(w http.ResponseWriter, r *http.Request) {

	t.todoService.handleCreate(w, r)
}

func (t *TodoController) updateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateHandler invoked")
	t.todoService.handleUpdate(w, r)
}

func (t *TodoController) deleteHandler(w http.ResponseWriter, r *http.Request) {
	t.todoService.handleDelete(w, r)
}
