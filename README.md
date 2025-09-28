# Базовая конфигурация для запуска контейнеров при помощи Docker Compose

<img alt="Docker Compose" src="https://raw.githubusercontent.com/ifdancoder/go-mysql-react-nginx-docker-compose/ac32d0d14ef0b4f66a0df63dd2205054dfbb6cea/static/media/docker-compose.png" height="350">

Основные конфигурационные файлы Docker Compose:
1. [DEV](./docker-compose.dev.yaml).
2. [PROD](./docker-compose.prod.yaml).

## MySQL

<img alt="MySQL" src="https://raw.githubusercontent.com/ifdancoder/go-mysql-react-nginx-docker-compose/ac32d0d14ef0b4f66a0df63dd2205054dfbb6cea/static/media/mysql.png" height="350">

Обыкновенный MySQL с двумя volume, один из которых необходим для хранения данных бд, а другой - для заполнения бд на основе дампов.

## Nginx

<img alt="Nginx" src="https://raw.githubusercontent.com/ifdancoder/go-mysql-react-nginx-docker-compose/ac32d0d14ef0b4f66a0df63dd2205054dfbb6cea/static/media/nginx.png" height="350">

Присутствует два конфигурационных файла (для DEV и PROD). Они оба заставляют Nginx проксировать запросы к /back/ как запросы для сервера на Go. Конфигурационный файл nginx определяется используемым конфигурационным файлом Docker Compose:
1. [DEV](./nginx/conf.d/dev.conf). Nginx проксирует запросы, кроме как по пути /back/, как запросы для тестового сервера на React.
2. [PROD](./nginx/conf.d/prod.conf). Nginx на всех путях, кроме /back/, обрабатывает папку build, полученную в результате билда React-приложения.

## Golang

<img alt="Golang" src="https://raw.githubusercontent.com/ifdancoder/go-mysql-react-nginx-docker-compose/ac32d0d14ef0b4f66a0df63dd2205054dfbb6cea/static/media/golang.png" height="350">

Предусмотрен build и запуск приложения на Go во время запуска контейнера (а не его билда). Билд и запуск происходит посредством файла [start.sh](./backend/start.sh).

## React

<img alt="React" src="https://raw.githubusercontent.com/ifdancoder/go-mysql-react-nginx-docker-compose/ac32d0d14ef0b4f66a0df63dd2205054dfbb6cea/static/media/react.png" height="350">

Предусмотрен build и запуск тестового сервера React, когда используется тестовая конфигурация Docker Compose и только build в противном случае.
