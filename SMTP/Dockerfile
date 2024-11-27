# SMTP/Dockerfile
FROM debian:bookworm-slim

# Installer Postfix, mailutils, telnet et netcat-openbsd
RUN apt-get update && \
    apt-get install -y postfix mailutils telnet netcat-openbsd && \
    echo "localhost" > /etc/mailname && \
    apt-get clean

# Copier la configuration Postfix
COPY main.cf /etc/postfix/

# Exposer le port SMTP
EXPOSE 25

# Démarrer Postfix en mode foreground
CMD ["postfix", "start-fg"]
