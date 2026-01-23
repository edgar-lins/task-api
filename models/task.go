package models

import "errors"

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// Validate verifica se a tarefa segue as regras de negócio
// Retorna nil se estiver tudo ok, ou um erro se falhar
func (t *Task) Validate() error {
	if t.Description == "" {
		return errors.New("a descrição não poder estar vazia")
	}
	if len(t.Description) < 3 {
		return errors.New("a descrição deve ter pelo menos 3 caracteres")
	}
	return nil
}
