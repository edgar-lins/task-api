package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"task-api/db"
	"task-api/models"
)

var database *sql.DB

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		// Chama a função da pasta db
		tasks, err := db.GetTasks(database)
		if err != nil {
			http.Error(w, "Errp ao buscar", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(tasks)

	case "POST":
		var newTask models.Task // Usa a struct do pacote models
		json.NewDecoder(r.Body).Decode(&newTask)

		// Chama a função da pasta db
		id, err := db.CreateTask(database, newTask)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		newTask.ID = id

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newTask)
	}
}

func main() {
	// 1. Inicializa o banco e guarda a conexão na variável global
	database = db.InitDB()
	// Importante: Fechar a conexão quando o main morrer (Ctrl+C)
	defer database.Close()

	// 2. Configura rotas
	http.HandleFunc("/tasks", tasksHandler)

	// 3. Sobe o servidor
	println("🔥 API Organizada rodando em http://localhost:8080/tasks")
	http.ListenAndServe(":8080", nil)
}
