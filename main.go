package main

import (
	"encoding/json"
	"net/http"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	// Simulando dados que viriam do banco de dados
	tasks := []Task{
		{ID: 1, Description: "Aprender HTTP", Done: true},
		{ID: 2, Description: "Criar primeira API", Done: false},
	}

	// Define que a resposta é do tipo JSON (importante para o navegador entender)
	w.Header().Set("Content-Type", "application/json")

	// Transforma a struct JSON e envia para o navegador (w)
	json.NewEncoder(w).Encode(tasks)
}

func main() {
	http.HandleFunc("/tasks", getTasks) // Mudamos a rota para /tasks

	println("Servidor rodando em http://localhost:8080/tasks")
	http.ListenAndServe(":8080", nil)
}
