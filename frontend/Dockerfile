# Frontend Dockerfile
FROM node:16 AS build
WORKDIR /app

COPY frontend/package*.json .
COPY .env /.env
RUN npm install

COPY frontend .
RUN npm run build

FROM nginx:alpine
COPY --from=build /app/build /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Containers run nginx with global directives and daemon off
ENTRYPOINT ["nginx", "-g", "daemon off;"]