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

// Variável GLOBAL: Agora a lista vive fora da função, para não resetar a cada requisição
var tasks = []Task{
	{ID: 1, Description: "Entender GET vs POST", Done: true},
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		// 🟢 Se for GET: Devolve a lista (igual fizemos antes)
		json.NewEncoder(w).Encode(tasks)
	case "POST":
		// 🔵 Se for POST: Cria tarefa nova
		var newTask Task

		// 1. Decodifica o JSON que veio no corpo da requisição (Body)
		// e joga para dentro da variável newTask
		err := json.NewDecoder(r.Body).Decode(&newTask)
		if err != nil {
			http.Error(w, "Erro ao ler o JSON", http.StatusBadRequest)
			return
		}

		// 2. Lógica simples de ID (pega o tamanho + 1)
		newTask.ID = len(tasks) + 1

		// 3. Adiciona na lista global
		tasks = append(tasks, newTask)

		// 4. Devolve o status 201 (Created) e a tarefa criada
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newTask)

	default:
		// Se tentarem DELETE ou PUT, devolvemos erro 405 (Método não permitido)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/tasks", tasksHandler) // Mudamos a rota para /tasks

	println("Servidor rodando em http://localhost:8080/tasks")
	http.ListenAndServe(":8080", nil)
}
