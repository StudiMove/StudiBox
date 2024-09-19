FROM node:18 AS build

WORKDIR /app

# Copier les fichiers de configuration
COPY package*.json ./
RUN npm install

# Copier tout le code source
COPY . .

# Construire le projet
RUN npm run build

# Étape 2 : Serveur
FROM nginx:alpine

# Copier les fichiers construits depuis l'étape de build
COPY --from=build /app/dist /usr/share/nginx/html

# Exposer le port 80
EXPOSE 80

# Démarrer Nginx
CMD ["nginx", "-g", "daemon off;"]
