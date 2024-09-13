<!-- my-app/
│
├── backend/
│ ├── main.go
│ ├── handlers/
│ ├── models/
│ ├── routes/
│ ├── utils/
│ ├── Dockerfile
│ ├── go.mod
│ ├── go.sum
│
├── frontend/
│ ├── src/
│ │ ├── components/
│ │ ├── pages/
│ │ ├── App.tsx
│ │ ├── index.tsx
│ ├── public/
│ ├── tsconfig.json
│ ├── tailwind.config.js
│ ├── package.json
│ ├── Dockerfile
│
├── docker-compose.yml
└── README.md


/backend -->

/backend
├── cmd/
│ └── server/
│ └── main.go # Point d'entrée principal du serveur
├── config/
│ └── config.go # Gestion de la configuration (DB, serveur, stockage)
├── internal/
│ ├── api/
│ │ ├── handlers/
│ │ │ ├── file/ # Dossier pour les gestionnaires de fichiers
│ │ │ │ ├── upload_file_handler.go # Gestionnaire pour le téléchargement des fichiers
│ │ │ │ ├── delete_file_handler.go # Gestionnaire pour la suppression des fichiers
│ │ │ │ └── get_file_url_handler.go # Gestionnaire pour récupérer l'URL des fichiers
│ │ │ ├── user/ # Dossier pour les gestionnaires d'utilisateurs
│ │ │ │ ├── user_handler.go # Gestionnaire pour les opérations sur les utilisateurs
│ │ │ │ └── auth_handler.go # Gestionnaire pour l'authentification
│ │ │ ├── authentication/ # Dossier pour l'authentification
│ │ │ │ ├── register_handler.go # Gestionnaire pour l'inscription
│ │ │ │ └── auth_handler.go # Gestionnaire pour la connexion
│ │ ├── routes/
│ │ │ ├── file_routes.go # Routes pour les fichiers
│ │ │ ├── user_routes.go # Routes pour les utilisateurs
│ │ │ └── auth_routes.go # Routes pour l'authentification
│ │ └── routes.go # Route principale
│ ├── services/
│ │ ├── storage/
│ │ │ ├── storage_service.go # Interface StorageService
│ │ │ ├── s3_storage.go # Implémentation pour Amazon S3
│ │ │ ├── cloudflare_storage.go # (Si jamais tu souhaites l'ajouter)
│ │ │ └── storage_errors.go # Gestion des erreurs de stockage
│ │ ├── user_service.go # Service pour les utilisateurs
│ │ └── ...
│ ├── db/
│ │ ├── migrations/
│ │ │ └── 001_initial_schema.go # Migrations pour créer la base de données
│ │ ├── models/
│ │ │ ├── user.go # Modèle utilisateur
│ │ │ └── transaction.go # Modèle transaction
│ │ └── repositories/
│ │ │ ├── user_repository.go # Repositories pour accéder aux données utilisateur
│ │ │ └── transaction_repository.go # Repositories pour accéder aux données transaction
│ ├── utils/
│ │ ├── logger.go # Gestionnaire de logs
│ │ └── security.go # Fonctions de sécurité (comme le hashage des mots de passe)
├── scripts/
│ ├── migrate.sh # Script pour les migrations de base de données
│ └── run.sh # Script pour exécuter l'application
├── test/
│ ├── integration/
│ │ └── user_integration_test.go # Tests d'intégration pour les utilisateurs
│ └── unit/
│ └── user_handler_test.go # Tests unitaires pour les gestionnaires d'utilisateurs
├── go.mod # Fichier de gestion des dépendances Go
└── go.sum # Fichier des sommes de contrôle des dépendances
