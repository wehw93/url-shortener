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
### 2. Установите зависимости:
go mod tidy  

