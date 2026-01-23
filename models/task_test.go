package models

import "testing"

func TestTaskValidation(t *testing.T) {

	// Cenário 1: Sucesso
	t.Run("Deve aceitar uma descrição válida", func(t *testing.T) {
		task := Task{Description: "Comprar leite"}
		err := task.Validate()

		if err != nil {
			t.Errorf("Esperava sucesso, mas recebeu erro: %v", err)
		}
	})

	// Cenário 2: Falha (Vazio)
	t.Run("Deve rejeitar descrição vazia", func(t *testing.T) {
		task := Task{Description: ""}
		err := task.Validate()

		if err == nil {
			t.Error("Esperava erro, mas a validação passou batido")
		}
	})

	// Cenário 3: Falha (Curto demais)
	t.Run("Deve rejeitar descrição muito curta", func(t *testing.T) {
		task := Task{Description: "Oi"}
		err := task.Validate()

		if err == nil {
			t.Error("Esperava erro para texto curto, mas passou")
		}

	})
}
