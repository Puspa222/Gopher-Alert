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
    ├── http/
    │   ├── module.go       
    │   ├── server.go       
    │   ├── handler.go
    │   └── middleware.go
    │   │
    │   ├── notifier/
    │   │   ├── interface.go
    │   │   ├── console.go
    │   │   ├── discord.go
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