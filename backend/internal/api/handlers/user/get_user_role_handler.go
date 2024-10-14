// backend/internal/api/handlers/user/get_user_role_handler.go
package user

import (
    "encoding/json"
    "log"
    "net/http"
    "strings"
    "fmt"
    "backend/config"
    "backend/internal/services/userservice" // Mise à jour du nom du package
    "backend/internal/utils"
)

// GetUserRoleHandler gère la récupération du rôle d'un utilisateur.
func GetUserRoleHandler(userService *userservice.UserService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Récupérer le token à partir des en-têtes
        tokenStr := r.Header.Get("Authorization")
        if tokenStr == "" {
            http.Error(w, "Unauthorized: missing token", http.StatusUnauthorized)
            return
        }

        // Vérifie que le token commence par "Bearer "
        if !strings.HasPrefix(tokenStr, "Bearer ") {
            http.Error(w, "Unauthorized: invalid token format", http.StatusUnauthorized)
            return
        }

        // Extraire le token sans le préfixe "Bearer "
        tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

        // Extraire l'ID de l'utilisateur à partir du token
        userID, err := utils.ExtractUserIDFromToken(tokenStr, config.AppConfig.JwtSecretAccessKey)
        if err != nil {
            log.Printf("Erreur d'extraction du token: %v", err)
            http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
            return
        }

        // Utiliser le UserService pour obtenir les rôles de l'utilisateur
        roles, err := userService.GetUserRolesByID(userID)
        if err != nil {
            http.Error(w, "User not found", http.StatusNotFound)
            return
        }

        // Extraire uniquement les noms des rôles
        var roleNames []string
        for _, role := range roles {
            roleNames = append(roleNames, role.Name)
        }

        // Renvoyer seulement le nom du rôle en réponse
        response := map[string]interface{}{
            "role": roleNames[0], // Si vous voulez juste le premier rôle (si un utilisateur a plusieurs rôles)
        }
        
        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(response); err != nil {
            log.Printf("Erreur lors de l'encodage de la réponse: %v", err)
            http.Error(w, fmt.Sprintf("Internal Server Error: %v", err), http.StatusInternalServerError)
        }
    }
}
