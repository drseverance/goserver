# Go API Server

A containerized REST API built with Go and PostgreSQL. This project demonstrates backend development, database integration, Docker deployment, health checks, and database migrations.

## Overview

This project is a lightweight REST API written in Go. It is containerized with Docker and deployed using Docker Compose on my homelab server running Ubuntu Server.

## Tech Stack

- Go 1.26.3
- PostgreSQL 17
- Docker
- Docker Compose
- REST API
- SQL migrations
- Ubuntu Server

## Features

- RESTful API endpoints
- PostgreSQL database integration
- Dockerized Go application
- Persistent PostgreSQL storage
- Database health checks
- API health endpoint
- Readiness endpoint
- Automated database migrations
- Environment-based configuration

## Project Structure

```text
goserver/
├── api/
│   ├── main.go
│   ├── Dockerfile
│   ├── go.mod
│   └── go.sum
│
├── migrations/
│   └── SQL migration files
│
├── docker-compose.yml
└── README.md
```
## Architecture

```text
                  Client
                    |
                    v
              Nginx Reverse Proxy
                    |
                    v
                Go REST API
                    |
                    v
              PostgreSQL Database
```

## Getting Started

### Prerequisites

- Docker
- Docker Compose
- Git

### Clone Repository

```bash
git clone https://github.com/drseverance/goserver.git
cd goserver
```

### Create Docker Network

Create the shared Docker network used by the reverse proxy.

```bash
docker network create public-network
```

If the network already exists, Docker will report that it already exists. This is expected and you can continue.

### Environment Configuration

Create a `.env` file in the project root.

```env
POSTGRES_HOST=postgres
POSTGRES_PORT=5432
POSTGRES_USER=your_username
POSTGRES_PASSWORD=your_password
POSTGRES_DB=your_database
```

### Start the Application

Build and start the containers.

```bash
docker compose up -d --build
```

Verify the containers are running.

```bash
docker ps
```

## API Endpoints

The API provides CRUD operations for managing users.

### Health Check

Check API and database connectivity:

```
GET /health
```

Example:

```bash
curl http://localhost:8991/health
```

Response:

```json
{
  "status": "ok"
}
```

---

## Users

### Get All Users

```
GET /users
```

Example:

```bash
curl http://localhost:8991/users
```

Response:

```json
[
  {
    "id": 1,
    "name": "Dustin",
    "email": "dustin@example.com"
  }
]
```

---

### Get User By ID

```
GET /users/{id}
```

Example:

```bash
curl http://localhost:8991/users/1
```

---

### Create User

```
POST /users
```

Example:

```bash
curl -X POST http://localhost:8991/users \
-H "Content-Type: application/json" \
-d '{"name":"New User","email":"new@example.com"}'
```

---

### Update User

```
PUT /users/{id}
```

Example:

```bash
curl -X PUT http://localhost:8991/users/1 \
-H "Content-Type: application/json" \
-d '{"name":"Updated User","email":"updated@example.com"}'
```

---

### Delete User

```
DELETE /users/{id}
```

Example:

```bash
curl -X DELETE http://localhost:8991/users/1
```
