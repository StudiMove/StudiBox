#!/bin/bash

# Se déplacer vers le dossier de déploiement
cd /path/to/your/deployment

# Arrêter les anciens conteneurs
docker-compose down

# Mettre à jour les images
docker-compose pull

# Démarrer les nouveaux conteneurs
docker-compose up -d
