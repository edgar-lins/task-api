package db

import (
	"database/sql"
	"log"
	"task-api/models" // Importamos a struct da pasta models

	_ "modernc.org/sqlite" // Importa o driver mas não usa diretamente (por isso o _)
)

func InitDB() *sql.DB {
	connection, err := sql.Open("sqlite", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT,
		done BOOLEAN
	);
	`
	_, err = connection.Exec(sqlStmt)
	if err != nil {
		log.Fatal("Erro ao criar tabela: %q", err)
	}

	return connection
}

// GetTasks busca as tarefas. Recebe a conexão como parâmetro.
func GetTasks(connection *sql.DB) ([]models.Task, error) {
	rows, err := connection.Query("SELECT id, description, done FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []models.Task{}
	for rows.Next() {
		var t models.Task
		// Note que agora usamos models.Task
		if err := rows.Scan(&t.ID, &t.Description, &t.Done); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

// CreateTask insere uma tarefa
func CreateTask(connection *sql.DB, task models.Task) (int, error) {
	result, err := connection.Exec("INSERT INTO tasks (description, done) VALUES (?, ?)", task.Description, task.Done)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	return int(id), nil
}
