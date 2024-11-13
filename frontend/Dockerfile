# Étape 1 : Build avec Node.js 20
FROM node:20 AS build

WORKDIR /app

# Copier et installer les dépendances avec `npm ci`
COPY package*.json ./
RUN npm ci

# Copier le code source et le construire
COPY . .
RUN npm run build

# Étape 2 : Serveur avec Nginx
FROM nginx:stable-alpine

# Copier les fichiers construits depuis l'étape de build
COPY --from=build /app/dist /usr/share/nginx/html

# Exposer le port 80 pour le serveur web
EXPOSE 80

# Démarrer Nginx
CMD ["nginx", "-g", "daemon off;"]
