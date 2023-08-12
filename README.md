# Reverse proxy c помощью go-chi middleware

Для решения этой задачи вам понадобится Docker и docker-compose.
OS Ubuntu:
```bash 
wget -O - https://gist.githubusercontent.com/ptflp/62f62201e8bd0683abdfbed631192db3/raw/docker-install.sh | bash
```

Установка на MacOS и Windows: [Docker install](https://docs.docker.com/get-docker/)

## Задача

Напишите middleware для go-chi, который будет перенаправлять запросы на другой сервер.
Если ресурс имеет префикс `/api/`, то запрос должен выдавать текст `Hello from API`
Все остальные запросы должны проксироваться на `http://hugo:1313` (сервер hugo). 

## Критерии приемки:

- Сервис proxy должен быть доступен на порту `8080`
- Сервис hugo должен быть доступен на порту `1313`
- Сервис hugo должен быть доступен по имени `hugo`
- Оба сервирса резолвятся через localhost (localhost:8080 и localhost:1313)

Доп материалы:
- [docker](https://go.ptflp.ru/course1/7/7.2/)