# Gopher-Alert: A Modular Notification Gateway

**Gopher-Alert** is a backend service for sending notifications to multiple providers (Console, Discord, Telegram, Slack, Email) via a single API call. Built for modularity and scalability, it showcases professional engineering practices: dependency injection, clean architecture, and UberFX lifecycle management.

---

## Table of Contents

- [Features](#features)
- [Project Structure](#project-structure)
- [Architecture](#architecture)
- [Getting Started](#getting-started)
- [API Usage](#api-usage)
- [Adding New Providers](#adding-new-providers)

---

## Features

- Modular notification system with provider interface
- UberFX for dependency injection and lifecycle management
- Gin HTTP API for sending notifications
- Supports multiple providers
- Easily extensible for future providers
- Clean separation of HTTP, service, and provider layers
- Optional API Key middleware for security

---

## Project Structure

```text
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
│   ├── middleware.go
│   ├── notifier/
│   │   ├── interface.go
│   │   ├── console.go
│   │   ├── discord.go
│   │   └── registry.go
│   ├── service/
│   │   └── notification_service.go
│   ├── storage/
│   │   └── sqlite.go
│   └── config/
│       └── config.go
│
├── go.mod
└── .env
```

---

## Architecture

```mermaid
flowchart TD
    A[Client (curl, Postman, frontend)]
    B[HTTP Layer (Gin)]
    C[Service Layer (NotificationService)]
    D[Provider Registry (Registry)]
    E[Notifiers (Console, Discord, etc.)]

    A --> B
    B --> C
    C --> D
    D --> E
```

- **HTTP Layer:** Handles requests, validation, and middleware.
- **Service Layer:** Contains the core business logic.
- **Notifier Layer:** Abstract interface for message sending. Individual providers implement this interface.
- **Registry:** Resolves provider names to concrete notifier instances.

---

## Getting Started

### Requirements

- Go 1.21+
- Git

### Install

```bash
git clone https://github.com/puspa222/gopher-alert.git
cd gopher-alert
go mod tidy
# Run the application
go run ./cmd/app
```

---
