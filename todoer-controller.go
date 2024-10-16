package main

import (
	"net/http"
)

type TodoController struct {
	todoService TodoService
}

func handleUnAuthorized(w http.ResponseWriter) {
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func (t *TodoController) getHandler(w http.ResponseWriter, r *http.Request) {
	if user, err := validateTokenFromUser(r); err == nil {
		t.todoService.handleGet(w, &user.Id)
	} else {
		handleUnAuthorized(w)
	}
}

func (t *TodoController) getByIdHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := validateTokenFromUser(r); err == nil {
		t.todoService.handleGetById(w, r)
	} else {
		handleUnAuthorized(w)
	}
}

func (t *TodoController) createHandler(w http.ResponseWriter, r *http.Request) {
	if user, err := validateTokenFromUser(r); err == nil {
		t.todoService.handleCreate(w, r, user)
	} else {
		handleUnAuthorized(w)
	}
}

func (t *TodoController) updateHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := validateTokenFromUser(r); err == nil {
		t.todoService.handleUpdate(w, r)
	} else {
		handleUnAuthorized(w)
	}
}

func (t *TodoController) deleteHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := validateTokenFromUser(r); err == nil {
		t.todoService.handleDelete(w, r)
	} else {
		handleUnAuthorized(w)
	}
}
