#!/bin/bash

# 1. Create the named volume for Redis
podman volume create redis_data

# 2. Build your local images
podman build -t backend:latest -f ./backend/Dockerfile ./backend
podman build -t frontend:latest -f ./frontend/ui/Dockerfile ./frontend/ui

# 3. Create the Pod and expose the ports
# (Notice we map 6069 and 6070 here)
podman pod create \
  --name personal-website-pod \
  --infra=true \
  -p 6069:6069 \
  -p 6070:6070

# 4. Start the Redis container inside the pod
podman run -d \
  --name redis-server \
  --pod personal-website-pod \
  --restart always \
  -v redis_data:/data:Z \
  docker.io/library/redis

# 5. Start the Backend container inside the pod
# NOTE: Update your .env.dev so that your Redis host config points to '127.0.0.1' or 'localhost'
podman run -d \
  --name backend \
  --pod personal-website-pod \
  --restart always \
  --env-file ./backend/.env.prod \
  backend:latest

# 6. Start the Frontend container inside the pod
# NOTE: Update your frontend config to point to 'localhost:6069' for backend API calls
podman run -d \
  --name frontend \
  --pod personal-website-pod \
  --restart always \
  frontend:latest
