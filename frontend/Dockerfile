# Frontend Dockerfile
FROM node:16 AS build
WORKDIR /app

COPY frontend/package*.json .

# NOTE: This placeholder URL will be replaced at entrypoint via environment variable
# LINK ./entrypoint.sh
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