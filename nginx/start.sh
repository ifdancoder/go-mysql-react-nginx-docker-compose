#!/bin/sh
set -e

# Подставляем переменные в конфиг
envsubst '$NGINX_HOST' < /etc/nginx/conf.d/default.conf.template > /etc/nginx/conf.d/default.conf

# Запускаем nginx
exec nginx -g 'daemon off;'