package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var getBaseSelectString = "SELECT todos.id, todo, done, createdDate, updateDate, user_id FROM todos FULL JOIN users ON todos.user_id=users.id"

type DBOperations struct {
	DB *sql.DB
}

func (d *DBOperations) saveToDb(todo Todo) error {
	_, err := d.DB.Exec("INSERT INTO todos (todo, done, active, user_id) VALUES ($1, $2, true, $3);",
		todo.Todo, todo.Done, todo.UserId)
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
	rows, err := d.DB.Query(getBaseSelectString+" WHERE id='$1' AND active=true;", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todo Todo
	for rows.Next() {
		if err := rows.Scan(&todo.Id, &todo.Todo, &todo.Done, &todo.CreatedDate, &todo.UpdateDate, &todo.UserId); err != nil {
			return nil, err
		}
	}
	return &todo, nil
}

func (d *DBOperations) getAllFromDb() ([]Todo, error) {
	rows, err := d.DB.Query(getBaseSelectString+" WHERE todos.active=$1 AND users.active=$1;", true)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.Id, &todo.Todo, &todo.Done, &todo.CreatedDate, &todo.UpdateDate, &todo.UserId); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (d *DBOperations) getUsersFromDb() ([]User, error) {
	rows, err := d.DB.Query("SELECT id, token, name FROM users WHERE active=$1;", true)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Token, &user.Name); err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (d *DBOperations) saveUserToDb(user User) error {
	_, err := d.DB.Exec("INSERT INTO users (token, name, active) VALUES ($1, $2, $3);",
		user.Token, user.Name, true)
	return err
}

func (d *DBOperations) updateUserInDb(user User) error {
	_, err := d.DB.Exec("UPDATE users SET name = '$2' WHERE token='$1';",
		user.Token, user.Name)
	return err
}
