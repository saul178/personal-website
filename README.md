# Personal Website Repo

# TODO
- [x] migrate from Docker to Podman
- [x] update this readme
- [x] begin the Home page
- [x] begin the Projects page
- [ ] begin the Personal page

- [x] look into transitions with vue (https://youtu.be/L77Uq93XXzk)
- [x] fix the alignment with the icon label -> LP
- [x] start the backend -> low priority
- [x] look into the github api for the projects page (https://github.com/google/go-github) -> LP
- [x] separate some of the code in the views files to be their own components like the projects cards and its tags -> LP

# Personal Portfolio Website

This repository contains the source code for my personal portfolio website.
> [!NOTE]
Website is currently down as I am in the process of moving it to a different host.

## Overview

The website is built with a modern full stack architecture. The frontend is developed using Vue.js and TailwindCSS, while the backend is implemented in Golang using the Gin web framework and Redis for caching.

The application is containerized with Docker and deployed on a bare metal server managed through Proxmox. Traffic is routed through Nginx Proxy Manager, and the domain is managed through Cloudflare.

## Tech Stack

### Frontend
- Vue.js
- TailwindCSS

### Backend
- Golang
- Gin (Go web framework)

### Infrastructure
- Docker
- Nginx
- Redis
- Nginx Proxy Manager
- Proxmox
- Cloudflare

## Deployment

The frontend and backend services are containerized using Docker and orchestrated with Docker Compose. The application is hosted on a self managed bare metal server running Proxmox. Nginx Proxy Manager handles reverse proxying and SSL management, while Cloudflare manages DNS and external traffic routing.
