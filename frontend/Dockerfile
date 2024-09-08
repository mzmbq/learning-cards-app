# Frontend Dockerfile
FROM node:16 AS build
WORKDIR /app

COPY frontend/package*.json .

# Set the environment variables
# ENV REACT_APP_BACKEND_URL_DEVELOPMENT=http://localhost:8080
ENV REACT_APP_BACKEND_URL_PRODUCTION=https://api.example.com

RUN npm install

COPY frontend .
RUN npm run build

FROM nginx:alpine
COPY --from=build /app/build /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

COPY frontend/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Containers run nginx with global directives and daemon off
ENTRYPOINT ["/entrypoint.sh"]