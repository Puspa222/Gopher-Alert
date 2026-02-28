# Gopher Alert


# Project Structure
```
    gopher-alert/
    │
    ├── cmd/
    │   └── app/
    │       └── main.go
    │
    ├── internal/
    │   ├── http/
    │   │   ├── handler.go
    │   │   └── middleware.go
    │   │
    │   ├── notifier/
    │   │   ├── interface.go
    │   │   ├── console.go
    │   │   ├── telegram.go
    │   │   └── registry.go
    │   │
    │   ├── service/
    │   │   └── notification_service.go
    │   │
    │   ├── storage/
    │   │   └── sqlite.go
    │   │
    │   └── config/
    │       └── config.go
    │
    ├── go.mod
    └── .env

```