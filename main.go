package main

import (
	"fmt"
	"net/http"
)

// Esse função é um "Handler". Ela lida com quem bate na porta do servidor.
// w = Writer (onde escrevemos a resposta para o usuário)
// r = Request (os dados que o usuário mandou, tipo IP, navegador, etc)
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem-vindo ao meu primeiro Backend em Go! 🚀")
	fmt.Println("Alguém acessou a home page!") // Isso aparece no seu terminal
}

func main() {
	// Roteamento: Quando alguém acessar "/", chame a função homePage
	http.HandleFunc("/", homePage)

	fmt.Println("Servidor rodando na porta 8080...")
	// Liga o servidor na porta 8080
	// O 'nil' significa que usaremos o roteador padrão do Go
	http.ListenAndServe(":8080", nil)
}
