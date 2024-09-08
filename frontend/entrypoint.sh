#!/bin/sh

# Replaces PLACEHOLDER_URL with the actual backend URL
# TODO: find better solution for this
sed -i "s|https://api.example.com|${REACT_APP_BACKEND_URL_PRODUCTION}|g" /usr/share/nginx/html/static/js/*.js
echo "Replaced placeholder with ${REACT_APP_BACKEND_URL_PRODUCTION}"

# Run nginx
nginx -g 'daemon off;'