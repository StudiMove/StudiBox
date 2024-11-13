# Étape 1 : Construction avec Go
FROM golang:1.23 AS builder

# Définir le répertoire de travail pour la phase de build
WORKDIR /app

# Copier les fichiers go.mod et go.sum, puis télécharger les dépendances
COPY go.mod go.sum ./
RUN go mod download

# Copier tout le code source dans le conteneur
COPY . .

# Construire l'application Go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/buildedApp ./main.go

# Étape 2 : Exécution avec Debian Slim
FROM debian:bookworm-slim

# Installer Node.js, MJML, et les certificats CA
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    ca-certificates \
    curl \
    npm && \
    npm install -g mjml && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Définir le répertoire de travail pour l'exécution
WORKDIR /app

# Copier l'exécutable depuis la phase de build
COPY --from=builder /app/buildedApp .

# Copier les templates MJML
COPY templates/ ./templates/

# Exposer le port sur lequel ton application écoute
EXPOSE 8080

# Commande pour exécuter l'application
CMD ["/app/buildedApp"]
