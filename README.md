В этом ДЗ вы создадите простейший RESTful CRUD.


Описание/Пошаговая инструкция выполнения домашнего задания:
Сделать простейший RESTful CRUD по созданию, удалению, просмотру и обновлению пользователей.
Пример API - https://app.swaggerhub.com/apis/otus55/users/1.0.0
Добавить базу данных для приложения.
Конфигурация приложения должна хранится в Configmaps.
Доступы к БД должны храниться в Secrets.
Первоначальные миграции должны быть оформлены в качестве Job-ы, если это требуется.
Ingress-ы должны также вести на url arch.homework/ (как и в прошлом задании)
На выходе должны быть предоставлена

ссылка на директорию в github, где находится директория с манифестами кубернетеса
инструкция по запуску приложения.
команда установки БД из helm, вместе с файлом values.yaml.
команда применения первоначальных миграций
команда kubectl apply -f, которая запускает в правильном порядке манифесты кубернетеса
Postman коллекция, в которой будут представлены примеры запросов к сервису на создание, получение, изменение и удаление пользователя. Важно: в postman коллекции использовать базовый url - arch.homework.
Проверить корректность работы приложения используя созданную коллекцию newman run коллекция_постман и приложить скриншот/вывод исполнения корректной работы

##Установка
1. kubectl apply -f ./manifests/manifests_db
2. helm install postgres -f ./manifests/values.yaml bitnami/postgresql --set volumePermissions.enabled=true
3. kubectl apply -f ./manifests/manifests_app

##Тестирование:
приложен скрин tests.png, а также сгенерированная коллкция postman



Список доступных методов `api`:
- POST http://arch.homework/api/user
- GET http://arch.homework/api/user/{userId}
- GET http://arch.homework/api/user/list
- PUT http://arch.homework/api/user/{userId}
- DELETE http://arch.homework/api/user/{userId}

Вспомогательные методы `api`:
- GET http://arch.homework/api/healthchecker
