package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // Importa o driver mas não usa diretamente (por isso o _)
)

var db *sql.DB

func initDB() {
	var err error
	// Abre (ou cria) o arquivo tasks.db
	db, err = sql.Open("sqlite", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	// Criamos a tabela se ela não existir (Comando SQL Puro!)
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT,
		done BOOLEAN
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal("Erro ao criar tabela: %q", err)
	}
}
