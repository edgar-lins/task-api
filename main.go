package main

import (
	"database/sql"
	"net/http"
	"task-api/db"
	"task-api/handlers"
)

var database *sql.DB

func main() {
	// 1. Inicializa o Banco
	database = db.InitDB()
	defer database.Close()

	// 2. Cria o Handler injetando a conexão do banco nele
	// Aqui estamos criando a "Caixa" e colocando o banco dentro
	myHandler := handlers.TaskHandler{
		DB: database,
	}

	// 3. Define as Rotas com Middleware
	// Note que chamamos myHandler.HandleTasks
	http.HandleFunc("/tasks", handlers.AuthMiddleware(myHandler.TasksHandler))

	println("🔥 API Organizada rodando em http://localhost:8080/tasks")
	http.ListenAndServe(":8080", nil)
}
