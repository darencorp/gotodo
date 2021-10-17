package todo

import (
	"encoding/json"
	"github.com/darencorp/gotodo/todo/models"
	utils "github.com/darencorp/gotodo/utils/http"
	"net/http"
	"time"
)

func init() {
	http.HandleFunc("/todo/all", utils.Get(getAll))
	http.HandleFunc("/todo/create", utils.Post(create))
	http.HandleFunc("/todo/get/", utils.Get(get))
	http.HandleFunc("/todo/update/", utils.Patch(update))
	http.HandleFunc("/todo/delete/", utils.Delete(remove))
}

func getAll(w http.ResponseWriter, _ *http.Request) {
	todos := models.GetAllTodos()
	todosData, err := json.Marshal(todos)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Unprocessable entity"))
		return
	}

	w.Write(todosData)
}

func create(w http.ResponseWriter, r *http.Request) {
	var newTodo models.Todo
	err := utils.ParseJsonRequest(w, r, &newTodo)

	if err != nil {
		return
	}

	if newTodo.Title == "" {
		utils.WriteBadRequest(w)
		return
	}

	now := time.Now()
	newTodo.CreatedAt = &now
	newTodo.Save()

	todoData, _ := json.Marshal(newTodo)
	w.Write(todoData)
}

func get(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetEntityId(w, r.URL)

	if err != nil {
		return
	}

	existedTodo := models.GetTodo(id)

	if existedTodo == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
		return
	}

	todoData, _ := json.Marshal(existedTodo)
	w.Write(todoData)
}

func update(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetEntityId(w, r.URL)

	if err != nil {
		return
	}

	var todoUpdate models.Todo
	err = utils.ParseJsonRequest(w, r, &todoUpdate)

	if err != nil {
		return
	}

	if todoUpdate.Title == "" {
		utils.WriteBadRequest(w)
		return
	}

	existedTodo := models.GetTodo(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
		return
	}

	todoUpdate.Id = existedTodo.Id
	todoUpdate.Save()

	todoData, _ := json.Marshal(todoUpdate)
	w.Write(todoData)
}

func remove(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetEntityId(w, r.URL)
	if err != nil {
		return
	}

	existedTodo := models.GetTodo(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
		return
	}

	existedTodo.Delete()
	w.WriteHeader(http.StatusNoContent)
}
