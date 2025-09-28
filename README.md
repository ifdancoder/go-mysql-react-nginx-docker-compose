# Базовая конфигурация для запуска контейнеров при помощи Docker Compose

## Docker Compose

<p align="center">
    <img alt="Docker Compose" src="https://raw.githubusercontent.com/ifdancoder/go-mysql-react-nginx-docker-compose/e4f708553ad0aa072b92cd22262d752338b37b9f/static/media/docker-compose.png" height="350">
</p>


Основные конфигурационные файлы Docker Compose:
1. [DEV](./docker-compose.dev.yaml).
2. [PROD](./docker-compose.prod.yaml).

Для Docker Compose необходим свой конфигурационный файл окружения .env ([Пример .env](./.env.example), новый .env необходимо разместить в той же директории, что и приведенный пример).

## Makefile

В репозитории представлен Makefile для удобства билда/запуска/выключения и прочих действий по отношению к сервисам Docker Compose. Чтобы увидеть возможные команды, добавленные при помощи Makefile, необходимо прописать:

```bash
make help
```

## Сервисы

Сервисы Docker Compose представлены и описаны ниже

### MySQL

<p align="center">
    <img alt="MySQL" src="https://raw.githubusercontent.com/ifdancoder/go-mysql-react-nginx-docker-compose/ac32d0d14ef0b4f66a0df63dd2205054dfbb6cea/static/media/mysql.png" height="350">
</p>


Обыкновенный MySQL с двумя volume, один из которых необходим для хранения данных бд, а другой - для заполнения бд на основе дампов [Директория для дампов](./db/dumps).

### Nginx

<p align="center">
    <img alt="Nginx" src="https://raw.githubusercontent.com/ifdancoder/go-mysql-react-nginx-docker-compose/ac32d0d14ef0b4f66a0df63dd2205054dfbb6cea/static/media/nginx.png" height="350">
</p>


Присутствует два конфигурационных файла (для DEV и PROD). Они оба заставляют Nginx проксировать запросы к `/back` как запросы для сервера на Go. Конфигурационный файл nginx определяется используемым конфигурационным файлом Docker Compose:
1. [DEV](./nginx/conf.d/dev.conf). Nginx проксирует запросы, кроме как по пути `/back`, как запросы для тестового сервера на React.
2. [PROD](./nginx/conf.d/prod.conf). Nginx на всех путях, кроме `/back`, обрабатывает папку build, полученную в результате билда React-приложения.
Билд и запуск происходит посредством файла [/nginx/start.sh](./nginx/start.sh).

### Golang

<p align="center">
    <img alt="Golang" src="https://raw.githubusercontent.com/ifdancoder/go-mysql-react-nginx-docker-compose/ac32d0d14ef0b4f66a0df63dd2205054dfbb6cea/static/media/golang.png" height="350">
</p>


Предусмотрен build и запуск приложения на Go во время запуска контейнера (а не его билда). Билд и запуск происходит посредством файла [/backend/start.sh](./backend/start.sh).

Для приложения на Go необходим свой конфигурационный файл окружения .env ([Пример .env](./backend/module/.env.example), новый .env необходимо разместить в той же директории, что и приведенный пример). 

### React

<p align="center">
    <img alt="React" src="https://raw.githubusercontent.com/ifdancoder/go-mysql-react-nginx-docker-compose/e4f708553ad0aa072b92cd22262d752338b37b9f/static/media/react.png" height="350">
</p>


Предусмотрен запуск приложения React, когда используется тестовая конфигурация Docker Compose и только build в противном случае.

Для приложения React необходим свой конфигурационный файл окружения .env ([Пример .env](./frontend/.env.example), новый .env необходимо разместить в той же директории, что и приведенный пример). 