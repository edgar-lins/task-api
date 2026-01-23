package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"task-api/db"
	"task-api/models"
)

// TaskHandler é uma struct que segura as dependências (o banco)
type TaskHandler struct {
	DB *sql.DB
}

func (h *TaskHandler) TasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		// Chama a função da pasta db
		tasks, err := db.GetTasks(h.DB)
		if err != nil {
			http.Error(w, "Erro ao buscar", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(tasks)

	case "POST":
		var newTask models.Task // Usa a struct do pacote models
		json.NewDecoder(r.Body).Decode(&newTask)

		// --- 🛑 NOVO: Validação ---
		if err := newTask.Validate(); err != nil {
			// Se der erro de validação, devolvemos erro 400 (Bad Request)
			// e a mensagem do erro (ex: "descrição muito curta")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// --------------------------

		// Chama a função da pasta db
		id, err := db.CreateTask(h.DB, newTask)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		newTask.ID = id

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newTask)
	}
}
