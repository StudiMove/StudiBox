package middleware

import (
    "context" // Ajoutez ceci si ce n'est pas déjà fait
    "net/http"
    "backend/internal/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }

        // Valider le JWT
        claims, err := utils.ValidateJWT(token, "your-secret-key")
        if err != nil {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }

        // Ajouter les claims au contexte
        ctx := context.WithValue(r.Context(), "user", claims)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
