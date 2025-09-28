# Базовая конфигурация для запуска контейнеров при помощи Docker Compose

<p align="center">
    <img alt="Docker Compose" src="https://raw.githubusercontent.com/ifdancoder/go-mysql-react-nginx-docker-compose/e4f708553ad0aa072b92cd22262d752338b37b9f/static/media/docker-compose.png" height="350">
</p>

Основные конфигурационные файлы Docker Compose:
1. [DEV](./docker-compose.dev.yaml).
2. [PROD](./docker-compose.prod.yaml).

## MySQL

<p align="center">
    <img alt="MySQL" src="https://raw.githubusercontent.com/ifdancoder/go-mysql-react-nginx-docker-compose/ac32d0d14ef0b4f66a0df63dd2205054dfbb6cea/static/media/mysql.png" height="350">
</p>

Обыкновенный MySQL с двумя volume, один из которых необходим для хранения данных бд, а другой - для заполнения бд на основе дампов.

## Nginx

<p align="center">
    <img alt="Nginx" src="https://raw.githubusercontent.com/ifdancoder/go-mysql-react-nginx-docker-compose/ac32d0d14ef0b4f66a0df63dd2205054dfbb6cea/static/media/nginx.png" height="350">
</p>

Присутствует два конфигурационных файла (для DEV и PROD). Они оба заставляют Nginx проксировать запросы к `/back` как запросы для сервера на Go. Конфигурационный файл nginx определяется используемым конфигурационным файлом Docker Compose:
1. [DEV](./nginx/conf.d/dev.conf). Nginx проксирует запросы, кроме как по пути `/back`, как запросы для тестового сервера на React.
2. [PROD](./nginx/conf.d/prod.conf). Nginx на всех путях, кроме `/back`, обрабатывает папку build, полученную в результате билда React-приложения.

Имеется конфигурационный файл ([Пример .env](./nginx/.env.example)) для удобства подстановки Nginx хоста в конфигурации. Билд и запуск происходит посредством файла [/nginx/start.sh](./nginx/start.sh).

## Golang

<p align="center">
    <img alt="Golang" src="https://raw.githubusercontent.com/ifdancoder/go-mysql-react-nginx-docker-compose/ac32d0d14ef0b4f66a0df63dd2205054dfbb6cea/static/media/golang.png" height="350">
</p>

Предусмотрен build и запуск приложения на Go во время запуска контейнера (а не его билда). Билд и запуск происходит посредством файла [/backend/start.sh](./backend/start.sh).

## React

<p align="center">
    <img alt="React" src="https://raw.githubusercontent.com/ifdancoder/go-mysql-react-nginx-docker-compose/e4f708553ad0aa072b92cd22262d752338b37b9f/static/media/react.png" height="350">
</p>

Предусмотрен build и запуск тестового сервера React, когда используется тестовая конфигурация Docker Compose и только build в противном случае.
