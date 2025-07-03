# Stage 1 - Build
FROM node:24 AS builder
WORKDIR /app

# Copy files into container
COPY . .

# Install deps and build
RUN npm install
RUN npm run build

# Stage 2 - Serve with nginx
FROM nginx:alpine

# Copy build output to nginx's html dir
COPY --from=builder /app/dist /usr/share/nginx/html

# Optional: custom nginx config
COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
