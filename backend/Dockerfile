# Étape 1 : Construction
FROM golang:1.23 AS builder

# Définir le répertoire de travail pour la phase de build
WORKDIR /app

# Copier les fichiers go.mod et go.sum puis télécharger les dépendances
COPY go.mod go.sum ./
RUN go mod download

# Copier tout le code source
COPY . .

# Construire l'application Go
RUN go build -o /app/buildedApp ./main.go

# Étape 2 : Exécution
FROM debian:bookworm-slim

# Installer les certificats CA et curl
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates curl && rm -rf /var/lib/apt/lists/*

# Définir le répertoire de travail pour l'exécution
WORKDIR /app

# Copier l'exécutable depuis la phase de build
COPY --from=builder /app/buildedApp .

# Exposer le port 8080
EXPOSE 8080

# Commande de démarrage de l'application
CMD ["/app/buildedApp"]
