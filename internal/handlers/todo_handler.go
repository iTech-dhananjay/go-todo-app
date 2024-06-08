package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"todo-app/internal/models"
	"todo-app/internal/services"
	"todo-app/pkg/responses"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoHandler struct {
	service *services.TodoService
}

func NewTodoHandler(db *models.DB) *TodoHandler {
	return &TodoHandler{
		service: services.NewTodoService(db),
	}
}

func (h *TodoHandler) HandleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.getAllTodos(w, r)
	case "POST":
		h.addTodo(w, r)
	default:
		responses.MethodNotAllowed(w)
	}
}

func (h *TodoHandler) HandleTodo(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/todos/")
	switch r.Method {
	case "GET":
		h.getTodoById(w, r, id)
	case "PUT":
		h.updateTodoById(w, r, id)
	case "DELETE":
		h.deleteTodoById(w, r, id)
	default:
		responses.MethodNotAllowed(w)
	}
}

func (h *TodoHandler) getAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.service.GetAll()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to get todos")
		return
	}
	responses.JSON(w, http.StatusOK, todos)
}

func (h *TodoHandler) addTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	todo.ID = primitive.NewObjectID().Hex()
	if err := h.service.Add(&todo); err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to add todo")
		return
	}
	responses.JSON(w, http.StatusCreated, todo)
}

func (h *TodoHandler) getTodoById(w http.ResponseWriter, r *http.Request, id string) {
	todo, err := h.service.GetById(id)
	if err != nil {
		responses.Error(w, http.StatusNotFound, "Todo not found")
		return
	}
	responses.JSON(w, http.StatusOK, todo)
}

func (h *TodoHandler) updateTodoById(w http.ResponseWriter, r *http.Request, id string) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := h.service.UpdateById(id, &todo); err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to update todo")
		return
	}
	responses.JSON(w, http.StatusOK, todo)
}

func (h *TodoHandler) deleteTodoById(w http.ResponseWriter, r *http.Request, id string) {
	if err := h.service.DeleteById(id); err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to delete todo")
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}
