// // backend/internal/api/routes/user_routes.go

package user_routes

import (
    "net/http"
    "backend/internal/api/handlers/user"
    userservice "backend/internal/services/userservice" // Import correct du service utilisateur
)

// RegisterUserRoutes enregistre toutes les routes liées aux utilisateurs
func RegisterUserRoutes(mux *http.ServeMux, userService *userservice.UserService) {
    // Enregistre la route pour récupérer les rôles des utilisateurs
    mux.HandleFunc("/user/role", user.GetUserRoleHandler(userService))
}

// backend/internal/api/routes/user_routes.go

// package user_routes

// import (
//     "net/http"
//     "backend/internal/api/handlers/user"
//     userservice "backend/internal/services/userservice"
//     "backend/internal/api/middlewares"  // Import du middleware
//     "backend/internal/services/auth"  // Import du service d'authentification pour les rôles
// )

// // RegisterUserRoutes enregistre toutes les routes liées aux utilisateurs
// func RegisterUserRoutes(mux *http.ServeMux, userService *userservice.UserService, authService *auth.AuthService) {
//     // Protéger la route avec le middleware d'authentification
//     mux.Handle("/user/role", middleware.AuthMiddleware(
//         http.HandlerFunc(user.GetUserRoleHandler(userService)),
//     ))

//     // Si tu veux protéger cette route par un rôle spécifique (par exemple "Admin"), tu peux ajouter RoleMiddleware
//     mux.Handle("/user/role/admin", middleware.AuthMiddleware(
//         middleware.RoleMiddleware(authService, "Admin", 
//             http.HandlerFunc(user.GetUserRoleHandler(userService)),
//         ),
//     ))
// }
