package main

import (
    "log"
    "net/http"
    "todo-app/internal/handlers"
    "todo-app/pkg/database"
)

func main() {
    db, err := database.Connect()
    if err != nil {
        log.Fatal(err)
    }

    handler := handlers.NewTodoHandler(db)

    http.HandleFunc("/todos", handler.HandleTodos)
    http.HandleFunc("/todos/", handler.HandleTodo)

    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
