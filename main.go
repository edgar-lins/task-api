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

// Note que removemos a variável global 'tasks'. O banco é a fonte da verdade agora.

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		// 🟢 SELECT: Busca as tarefas no banco
		rows, err := db.Query("SELECT id, description, done FROM tasks")
		if err != nil {
			http.Error(w, "Erro ao buscar tarefa", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		tasks := []Task{}
		for rows.Next() {
			var t Task
			// Scaneia cada linha do banco para dentro da struct
			rows.Scan(&t.ID, &t.Description, &t.Done)
			tasks = append(tasks, t)
		}
		json.NewEncoder(w).Encode(tasks)

	case "POST":
		// 🔵 INSERT: Salva no banco
		var newTask Task
		json.NewDecoder(r.Body).Decode(&newTask)

		// Executa o comando SQL de inserção
		// O "?" é um placeholder de segurança (evita SQL Injection)
		result, err := db.Exec("INSERT INTO tasks (description, done) VALUES (?, ?)", newTask.Description, newTask.Done)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Pega o ID que acabou de ser gerado pelo banco
		id, _ := result.LastInsertId()
		newTask.ID = int(id)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newTask)

	default:
		// Se tentarem DELETE ou PUT, devolvemos erro 405 (Método não permitido)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	// Inicializa o banco de dados antes de ligar o servidor
	initDB()

	http.HandleFunc("/tasks", tasksHandler)
	println("Servidor rodando em http://localhost:8080/tasks")
	http.ListenAndServe(":8080", nil)
}
