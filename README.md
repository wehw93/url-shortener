# URL-Shortener: Быстрый и надежный сервис для сокращения URL 🚀  

**URL-Shortener** — это мощный REST API-сервис для управления короткими ссылками. С его помощью вы можете быстро создавать, удалять и редиректить короткие ссылки. Проект разработан на **Golang**, что гарантирует высокую производительность и масштабируемость.  

## 🔧 Основные особенности  
- **Современные технологии**: Используется легковесный HTTP-роутер `go-chi/chi`, обеспечивающий гибкость и скорость.  
- **Поддержка баз данных**:  
  - **SQLite**: Идеально подходит для локальной разработки или легких задач.  
  - **MongoDB**: Подходит для распределенных и масштабируемых систем.  
- **Docker-Compose**: Упрощенная настройка окружения с помощью Docker Compose.  
- **Файл конфигурации**: Поддержка YAML для удобной настройки сервиса.  
- **Логирование**: Интеграция с `slog` для прозрачного отслеживания работы.  
- **Тестирование**: Ключевые компоненты покрыты тестами, что гарантирует надежность в продакшене.  

## 📜 API Обзор  
- **POST /url**: Создать короткую ссылку.  
- **GET /{alias}**: Редирект на оригинальную ссылку.  
- **DELETE /url/{alias}**: Удалить короткую ссылку.  

## 🗂️ Структура проекта  
```plaintext
├── cmd
│   └── url-shortener
│       └── main.go
├── config
│   └── local.yaml
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   ├── http-server
│   │   ├── handlers
│   │   │   └── url
│   │   │       ├── deleteURL
│   │   │       │   └── deleteURL.go
│   │   │       ├── redirect
│   │   │       │   └── redirect.go
│   │   │       └── save
│   │   │           └── save.go
│   │   └── middleware
│   │       └── logger
│   │           └── logger.go
│   ├── lib
│   │   ├── api
│   │   │   ├── api.go
│   │   │   └── response
│   │   │       └── response.go
│   │   ├── logger
│   │   │   └── sl
│   │   │       └── sl.go
│   │   └── random
│   │       ├── random.go
│   │       └── random_test.go
│   └── storage
│       ├── mongo
│       │   └── mongo.go
│       ├── sqlite
│       │   └── sqlite.go
│       └── storage.go
├── local.env
├── Makefile
├── README.md
├── storage.db
├── tests
│   └── url_shortener_test.go
└── url-shortener
```
##  Установка и запуск  

Чтобы установить и запустить **URL-Shortener**, выполните следующие шаги:  

### 1. Клонируйте репозиторий:  
```bash
git clone git@github.com:your-username/url-shortener.git  
cd url-shortener
```  
### 2. Установите зависимости:
```
go mod tidy
```
### 3. Настройте конфигурацию:
 - Используйте файл ```config/local.yaml``` для конфигурации вашего запуска.
 - Конфиг файл по умолчанию:
```
env: "local" # Возможные значения: local, dev, prod
http_server:
  address: "localhost:8080"
  timeout: 4s
  idle_timeout: 60s
  user: "myuser"
  password: "mypass"
storage:
  type: "SQLite" # Возможные значения: SQLite, MongoDB
  SQLite:
    path: "./storage.db"
  MongoDB:
    uri: "mongodb://admin:my_password@localhost:27017/url-shortener"
```
### 4. Запустите серверЖ
```
go run cmd/url-shortener/main.go  
```
### 5. Проверьте работу API:
- Вы можете использовать curl или Postman для тестирования:
- Создание короткой ссылки
  ```
  curl -X POST http://localhost:8080/url -u myuser:mypass -d '{"url": "https://github.com/wehw93", "alias": "my_github"}'
  ```
- Редирект по короткой ссылке:
  ```
  curl -X GET http://localhost:8080/my_github -u myuser:mypass
  ```
- Удаление короткой ссылки:
  ```
  curl -X DELETE http://localhost:8080/url/my_github -u myuser:mypass
  ```
  

