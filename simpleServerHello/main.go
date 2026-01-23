package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
	fmt.Println("Alguém está acessando a API")

}

func main() {
	http.HandleFunc("/hello", sayHello)

	fmt.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", nil)
}
