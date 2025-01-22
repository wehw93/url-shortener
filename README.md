# URL-Shortener: A Fast and Reliable URL Shortening Service 🚀  

**URL-Shortener** is a powerful REST API service for managing short URLs. With this service, you can quickly create, delete, and redirect short links. Built with **Golang**, it ensures high performance and scalability.  

## 🔧 Features  
- **Modern Technologies**: Utilizes the lightweight HTTP router `go-chi/chi` for flexible and efficient routing.  
- **Database Support**:  
  - **SQLite**: Perfect for local development or lightweight use cases.  
  - **MongoDB**: Ideal for distributed and scalable systems.  
- **Docker-Compose**: Simplified environment setup using Docker Compose.  
- **Config File**: YAML configuration support for easy and customizable settings.  
- **Logging**: Integrated with `slog` for transparent activity tracking.  
- **Thoroughly Tested**: All key components are covered with tests, ensuring reliability and production readiness.  

## 📜 API Overview  
- **POST /url**: Create a short URL.  
- **GET /{alias}**: Redirect to the original URL.  
- **DELETE /url/{alias}**: Delete a short URL.  
