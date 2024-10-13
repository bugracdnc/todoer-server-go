package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

type TodoRepository struct {
	todos []Todo
	DB    DBOperations
}

func (r *TodoRepository) getAll() *[]Todo {
	fetchedTodos, err := r.DB.getAllFromDb()
	if err != nil {
		log.Fatalf("Failed to get all from database: %v", err)
	}
	r.todos = fetchedTodos
	return &r.todos
}

func (r *TodoRepository) getById(id uuid.UUID) *Todo {
	todo, err := r.DB.getByIdFromDb(id)
	if err != nil {
		log.Fatalf("Failed to get by id from database: %v", err)
	}
	return todo
}

func (r *TodoRepository) save(todoToSave Todo) {
	fmt.Println(todoToSave.Todo)
	if err := r.DB.saveToDb(todoToSave); err != nil {
		log.Fatalf("Failed to save to database: %v", err)
	}
}

func (r *TodoRepository) update(todo Todo) {
	if err := r.DB.updateDoneInDb(todo); err != nil {
		log.Fatalf("Failed to update database: %v", err)
	}
}

func (r *TodoRepository) delete(id uuid.UUID) {
	if err := r.DB.deactivateInDb(id); err != nil {
		log.Fatalf("Failed to update database: %v", err)
	}
}
