package main

import (
	"database/sql"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type DBOperations struct {
	DB *sql.DB
}

func (d *DBOperations) saveToDb(todo Todo) error {
	_, err := d.DB.Exec("INSERT INTO todos (todo, done, createdDate, updateDate, active) VALUES ($1, $2, $3, $4, true);",
		todo.Todo, todo.Done, todo.CreatedDate, todo.UpdateDate)
	return err
}

func (d *DBOperations) updateDoneInDb(todo Todo) error {
	_, err := d.DB.Exec("UPDATE todos SET done = '$1' WHERE id = '$2';", todo.Done, todo.Id)
	return err
}

func (d *DBOperations) deactivateInDb(id uuid.UUID) error {
	_, err := d.DB.Exec("UPDATE todos SET active = false WHERE id = '$1';", id)
	return err
}

func (d *DBOperations) getByIdFromDb(id uuid.UUID) (*Todo, error) {
	rows, err := d.DB.Query("SELECT * FROM todos WHERE id='$1';", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todo Todo
	for rows.Next() {
		if err := rows.Scan(&todo.Id, &todo.Todo, &todo.Done, &todo.CreatedDate, &todo.UpdateDate); err != nil {
			return nil, err
		}
	}
	return &todo, nil
}

func (d *DBOperations) getAllFromDb() ([]Todo, error) {
	rows, err := d.DB.Query("SELECT id, todo, done, createdDate, updateDate FROM todos;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.Id, &todo.Todo, &todo.Done, &todo.CreatedDate, &todo.UpdateDate); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
