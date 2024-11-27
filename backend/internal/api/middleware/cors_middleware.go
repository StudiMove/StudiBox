// backend/internal/api/middleware/cors_middleware.go
package middleware

import (
	"log"
	"net/http"
)

// CORSMiddleware gère les requêtes CORS
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received %s request for %s", r.Method, r.URL.Path)

		// Définir les en-têtes CORS
		w.Header().Set("Access-Control-Allow-Origin", "https://studibox.fr")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin, X-Auth-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "1800")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Si la méthode est OPTIONS, renvoyer directement une réponse 200 OK
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Passer la requête au prochain handler
		next.ServeHTTP(w, r)
	})
}
