package main

import (
	"fmt"

	"github.com/google/uuid"
)

type TodoRepository struct {
	todo []Todo
}

func (r *TodoRepository) getAll() *[]Todo {
	return &r.todo
}

func (r *TodoRepository) getById(id uuid.UUID) *Todo {
	for _, item := range r.todo {
		if id == item.Id {
			return &item
		}
	}
	return nil
}

func (r *TodoRepository) save(todoToSave Todo) *Todo {
	r.todo = append(r.todo, todoToSave)
	return r.getById(todoToSave.Id)
}

func (r *TodoRepository) update(todoToSave Todo) *Todo {
	todoFetched := r.getById(todoToSave.Id)
	fmt.Println(todoToSave.Done)
	todoFetched.Done = todoToSave.Done
	todoFetched.Todo = todoToSave.Todo
	todoFetched.UpdateDate = todoToSave.UpdateDate
	for index, todo := range r.todo {
		if todo.Id == todoToSave.Id {
			r.todo[index] = todoToSave
			return &r.todo[index]
		}
	}
	return nil
}

func (r *TodoRepository) delete(id uuid.UUID) bool {
	deleted := false
	for index, todo := range r.todo {
		if todo.Id == id {
			r.todo = append(r.todo[:index], r.todo[index+1:]...)
			deleted = true
			break
		}
	}
	return deleted
}
