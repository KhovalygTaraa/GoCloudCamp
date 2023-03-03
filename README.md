# GoCloudCamp
# 1. Вопросы для разогрева
1. Опишите самую интересную задачу в программировании, которую вам приходилось решать?
+ *Реализация трассировки лучей. Только холст и метод putPixel(). Реализация STL языка С++. Познакомился с generics, самым сложным была реализация красно-черного дерева.*
2. Расскажите о своем самом большом факапе? Что вы предприняли для решения проблемы?
+ *Удалил весь учебный проект на финальном этапе командой rm -rf. Теперь всегда пушу в репозиторий*
3. Каковы ваши ожидания от участия в буткемпе?
+ *Дальнейший переход в штат*

# Часть 1. Разработка музыкального плейлиста

Выполнено в папке core

# Часть 2: Построение API для музыкального плейлиста

1. Выполнено в папке service.
2. Подключение к бд и listening находятся в пакете main.
3. Использован grpc, proto file находится в папке proto, собранный пакет находится в папке api.
3. Написан Dockerfile для сервиса.
4. Написан docker-compose, поднимается postgres и приложение music_service.
5. Для поднятия сервисов использовать команду make docker
