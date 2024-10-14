package middleware

import (
    "context" // Assurez-vous que ceci est importé
    "backend/config" // Importez le package où se trouve votre configuration
    "net/http"
    "log"
    "backend/internal/utils"
)

// func AuthMiddleware(next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         token := r.Header.Get("Authorization")
//         if token == "" {
//             http.Error(w, "Forbidden", http.StatusForbidden)
//             return
//         }

//         // Valider le JWT
//         claims, err := utils.ValidateJWT(token, config.AppConfig.JwtSecretAccessKey)
//         if err != nil {
//             http.Error(w, "Forbidden", http.StatusForbidden)
//             return
//         }

//         // Ajouter les claims au contexte
//         ctx := context.WithValue(r.Context(), "user", claims)
//         next.ServeHTTP(w, r.WithContext(ctx))
//     })
// }
// func AuthMiddleware(next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         token := r.Header.Get("Authorization")
//         if token == "" {
//             log.Println("No token found in the request")
//             http.Error(w, "Forbidden", http.StatusForbidden)
//             return
//         }

//         // Valider le JWT
//         claims, err := utils.ValidateJWT(token, config.AppConfig.JwtSecretAccessKey)
//         if err != nil {
//             log.Printf("Failed to validate token: %v", err)
//             http.Error(w, "Forbidden", http.StatusForbidden)
//             return
//         }

//         // Ajouter les claims au contexte
//         log.Printf("Token validated, user ID: %v", claims.UserID)
//         ctx := context.WithValue(r.Context(), "user", claims)
//         next.ServeHTTP(w, r.WithContext(ctx))
//     })
// }

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            log.Println("No token found in the request")

            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }

        // Supprimer "Bearer " du token
        if len(token) > 7 && token[:7] == "Bearer " {
            token = token[7:] // On retire "Bearer " pour ne garder que le token
        } else {
            http.Error(w, "Invalid token format", http.StatusUnauthorized)
            return
        }

        // Valider le JWT
        claims, err := utils.ValidateJWT(token, config.AppConfig.JwtSecretAccessKey)
        if err != nil {
            log.Printf("Failed to validate token: %v", err)

            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }

        // Ajouter les claims au contexte
        log.Printf("Token validated, user ID: %v", claims.UserID)

        ctx := context.WithValue(r.Context(), "user", claims)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
