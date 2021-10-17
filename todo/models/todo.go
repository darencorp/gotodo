package models

import (
	"github.com/darencorp/gotodo/sql"
	"time"
)

type Todo struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func (t *Todo) Save() {
	if t.Status == "" {
		t.Status = "TODO"
	}

	pgTodo := sql.PgTodo{
		Id:          t.Id,
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		CreatedAt:   t.CreatedAt,
		DeletedAt:   t.DeletedAt,
	}

	var savedTodo sql.PgTodo

	if t.Id == 0 {
		savedTodo = sql.CreateTodo(&pgTodo)
	} else {
		savedTodo = sql.UpdateTodo(&pgTodo)
	}

	t.Id = savedTodo.Id
	t.Title = savedTodo.Title
	t.Description = savedTodo.Description
	t.Status = savedTodo.Status
}

func (t *Todo) Delete() {
	sql.DeleteTodo(t.Id)
}

func GetAllTodos() []Todo {
	var todos []Todo
	fetchedTodos := sql.GetAllTodos()

	for _, fetchedTodo := range fetchedTodos {
		todos = append(todos, Todo{
			Id:          fetchedTodo.Id,
			Title:       fetchedTodo.Title,
			Description: fetchedTodo.Description,
			Status:      fetchedTodo.Status,
			CreatedAt:   fetchedTodo.CreatedAt,
			DeletedAt:   fetchedTodo.DeletedAt,
		})
	}

	if len(todos) != 0 {
		return todos
	}

	return make([]Todo, 0)
}

func GetTodo(id int) *Todo {
	fetchedTodo := sql.GetTodo(id)

	if fetchedTodo != nil {
		todo := &Todo{
			Id:          fetchedTodo.Id,
			Title:       fetchedTodo.Title,
			Description: fetchedTodo.Description,
			Status:      fetchedTodo.Status,
			CreatedAt:   fetchedTodo.CreatedAt,
			DeletedAt:   fetchedTodo.DeletedAt,
		}

		return todo
	}

	return nil
}
