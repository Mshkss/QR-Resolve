(ENG/RU)
QR-Resolve

-eng:
A service for processing QR codes containing URLs of the format https://example.com/resolve/{category}_{device_id} (e.g., https://example.com/resolve/a_12345678).

How it works:
A QR code is printed on a product, embedding a URL that points to this service.
The server receives the request, parses the URL to extract category (e.g., a) and device_id (e.g., 12345678).
Using the device_id, the server queries a MongoDB database, which stores a dictionary mapping device_id to one or more media service links (e.g., youtube.com).
The server returns the associated media link(s).
Database structure:
A MongoDB collection where the key is device_id (string) and the value is one or more links to media services (e.g., youtube.com/watch?v=...).

API:

The service exposes an API built with the Echo framework (Go).
Endpoint: GET /resolve/{category}_{device_id}.
Example request: https://example.com/resolve/a_12345678.
Response: JSON containing a media_link field (or an array of links).

Technologies:
Echo: A Go web framework for handling HTTP requests.
OpenAPI Codegen: Used for generating API specifications and client code.
MongoDB Driver: For interacting with the MongoDB database.

-ru:
Сервис для обработки QR-кодов, содержащих ссылки вида https://example.com/resolve/{category}_{device_id} (например, https://example.com/resolve/a_12345678).

Описание работы:

На товаре размещён QR-код, в котором закодирована ссылка на запрос к сервису.
Сервер принимает запрос, извлекает из ссылки параметры category (например, a) и device_id (например, 12345678).
На основе device_id сервер обращается к базе данных MongoDB, где хранится словарь вида {device_id: media_link}.
Сервер возвращает ссылку (или ссылки) на медиасервис (например, youtube.com), связанную с device_id.

Структура базы данных:
Коллекция в MongoDB, где ключом является device_id (строка), а значением — одна или несколько ссылок на медиасервисы (например, youtube.com/watch?v=...).

API:
Сервис предоставляет API, реализованное с использованием фреймворка Echo (Go).
Эндпоинт: GET /resolve/{category}_{device_id}.
Пример запроса: https://example.com/resolve/a_12345678.
Ответ: JSON с полем media_link (или массивом ссылок).

Технологии:
Echo: веб-фреймворк для Go, используется для обработки HTTP-запросов.
OpenAPI Codegen: генерация спецификации API и клиентского кода.
MongoDB Driver: взаимодействие с базой данных MongoDB.
