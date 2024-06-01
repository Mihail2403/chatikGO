# AUTH Service

## start:

- docker-compose up -d

## stop:

- docker-compose down

## api-url:

    http://localhost:8000/auth

## docs:

    http://localhost:8000/swagger/index.html

###### Если вылетает ошибка, добавь в /etc/docker/daemon.json следующее:

###### { "dns": ["192.168.2.1", "8.8.8.8", "8.8.8.4"] }
