package sql

import (
	"context"
	"time"
)

type PgTodo struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func GetAllTodos() []PgTodo {
	conn, err := pool.Acquire(context.Background())
	defer conn.Release()

	if err != nil {
		panic("Cannot acquire connection from pool")
	}

	var todos []PgTodo

	rows, _ := conn.Query(context.TODO(), "select id, title, description, status, created_at, deleted_at from todos;")
	defer rows.Close()

	for rows.Next() {
		var todo PgTodo
		err = rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Status, &todo.CreatedAt, &todo.DeletedAt)

		if err != nil {
			panic("Cannot read todo data from db.")
		}

		todos = append(todos, todo)
	}

	return todos
}

func GetTodo(id int) *PgTodo {
	conn, err := pool.Acquire(context.Background())
	defer conn.Release()

	if err != nil {
		panic("Cannot acquire connection from pool")
	}

	var todo PgTodo
	err = conn.QueryRow(context.TODO(), "select id, title, description, status, created_at, deleted_at from todos where id = $1;", id).Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Status, &todo.CreatedAt, &todo.DeletedAt)

	if err != nil {
		return nil
	}

	return &todo
}

func CreateTodo(todo *PgTodo) PgTodo {
	conn, err := pool.Acquire(context.Background())
	defer conn.Release()

	if err != nil {
		panic("Cannot acquire connection from pool")
	}

	var newTodo PgTodo
	err = conn.QueryRow(context.TODO(), "insert into todos(title, description, status, created_at, deleted_at) values($1, $2, $3, $4, $5) returning id, title, description, status, created_at, deleted_at;", todo.Title, todo.Description, todo.Status, todo.CreatedAt, todo.DeletedAt).Scan(&newTodo.Id, &newTodo.Title, &newTodo.Description, &newTodo.Status, &newTodo.CreatedAt, &newTodo.DeletedAt)

	if err != nil {
		panic("There was an error with todo data")
	}

	return newTodo
}

func UpdateTodo(todo *PgTodo) PgTodo {
	conn, err := pool.Acquire(context.Background())
	defer conn.Release()

	if err != nil {
		panic("Cannot acquire connection from pool")
	}

	var updatedTodo PgTodo
	err = conn.QueryRow(context.TODO(), "update todos set title=$1, description=$2, status=$3 where id=$4 returning id, title, description, status, created_at, deleted_at;", todo.Title, todo.Description, todo.Status, todo.Id).Scan(&updatedTodo.Id, &updatedTodo.Title, &updatedTodo.Description, &updatedTodo.Status, &updatedTodo.CreatedAt, &updatedTodo.DeletedAt)

	if err != nil {
		panic("There was an error with todo data")
	}

	return updatedTodo
}

func DeleteTodo(id int) {
	conn, err := pool.Acquire(context.Background())
	defer conn.Release()

	if err != nil {
		panic("Cannot acquire connection from pool")
	}

	commandTag, _ := conn.Exec(context.TODO(), "delete from todos where id=$1", id)

	if commandTag.RowsAffected() != 1 {
		panic("No todo found to delete")
	}
}
