package handlers

import "net/http"

// AuthMiddleware verifica o token no Header
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-API-Token")

		// Num cenário real, validaríamos contra o banco ou variáveis de ambiente
		if token != "minha-senha-secreta-123" {
			http.Error(w, "⛔️ Acesso Negado: Token inválido ou ausente", http.StatusUnauthorized)
			return
		}

		// Se passou, segue o jogo
		next(w, r)
	}
}
