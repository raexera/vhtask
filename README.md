# VHTask

This is a simple To-Do app.

## Features

- Create Task
- List Tasks
- Get Task by ID
- Update Task
- Delete Task
- Swagger Documentation

## Tech Stack

- Go
- Echo
- PostgreSQL
- Docker

## Project Structure:

```text
.
├── cmd
│   └── main.go
├── docker-compose.yml
├── Dockerfile
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── internal
│   ├── application
│   │   └── task_service.go
│   ├── domain
│   │   └── task.go
│   ├── infrastructure
│   │   └── task_repository.go
│   └── interface
│       └── task_handler.go
└── README.md
```

## Getting Started

1. Clone the Repository:

```sh
git clone https://github.com/raexera/vhtask.git
cd vhtask
```

2. Create and Configure .env files:

```sh
cp ./.env.example ./.env
```

3. Run with Docker Compose:

```sh
docker compose up --build
```

4. Access the Application: The API will be running at `http://localhost:8080`.

5. Swagger API Documentation: Visit `http://localhost:8080/swagger/index.html` for API documentation and testing.
