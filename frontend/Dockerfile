# Étape 1 : Build
FROM node:20-alpine AS build

# Définir le répertoire de travail
WORKDIR /app

# Copier uniquement les fichiers nécessaires pour installer les dépendances
COPY package*.json ./

# Installer uniquement les dépendances de production avec une gestion des peer-dependencies
RUN npm install --legacy-peer-deps

# Copier tout le code source
COPY . .

# Construire le projet pour la production
RUN npm run build

# Étape 2 : Serveur Nginx
FROM nginx:alpine

# Supprime les fichiers temporaires pour économiser de l'espace
RUN rm -rf /var/cache/apk/*

# Copier les fichiers construits depuis l'étape de build
COPY --from=build /app/dist /usr/share/nginx/html

# Exposer le port 80 pour le serveur web
EXPOSE 80

# Lancer Nginx
CMD ["nginx", "-g", "daemon off;"]
