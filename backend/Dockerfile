# Étape 1 : Construction
FROM golang:1.23 AS builder

# Définir le répertoire de travail
WORKDIR /app

# Copier go.mod et go.sum et télécharger les dépendances
COPY go.mod go.sum ./
RUN go mod download

# Copier le code source
COPY . .

# Construire l'application
WORKDIR /app/cmd/server
RUN go build -o main .

# Vérifiez que le binaire a bien été créé
RUN ls -la /app/cmd/server
RUN pwd

# Étape 2 : Exécution
FROM debian:bookworm-slim

# Installer les certificats CA
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*

# Définir le répertoire de travail
WORKDIR /app

# Copier l'exécutable de l'étape de construction
COPY --from=builder /app/cmd/server/main .

# Vérifiez le contenu du répertoire avant de démarrer
RUN ls -la /app
RUN pwd

# Copier le fichier .env dans l'image
COPY .env ./.env

# Exposer le port
EXPOSE 8080

# Définir le point d'entrée
CMD ["/app/main"]
