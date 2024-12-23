FROM node:20-alpine AS build

WORKDIR /app

# Copier uniquement les fichiers nécessaires pour installer les dépendances
COPY package*.json ./

# Installer uniquement les dépendances de production avec une gestion des peer-dependencies
RUN npm install --legacy-peer-deps

# Copier tout le code source
COPY . .

# Construire le projet
RUN npm run build

# Étape 2 : Serveur
FROM nginx:alpine

# Supprime les fichiers temporaires pour économiser de l'espace
RUN rm -rf /var/cache/apk/*

# Copier les fichiers construits depuis l'étape de build
COPY --from=build /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Exposer le port 80
EXPOSE 80

# Démarrer Nginx
CMD ["nginx", "-g", "daemon off;"]