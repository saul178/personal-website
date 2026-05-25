# build stage
FROM node:24-alpine AS build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .

RUN npm run build

# production stage
FROM docker.io/library/nginx:alpine AS production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
COPY nginx-prod.conf /etc/nginx/nginx.conf
# COPY nginx-dev.conf /etc/nginx/nginx.conf
EXPOSE 6070
CMD ["nginx", "-g", "daemon off;"]
