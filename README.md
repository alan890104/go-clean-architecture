# Go Clean Architecture (Library)

## Description

This repo aims to implement the Clean Architecture in Go by an library example. This project is inspired by the [go-clean-arch](https://github.com/bxcodec/go-clean-arch) and [go-backend-clean-architecture](https://amitshekhar.me/blog/go-backend-clean-architecture). The implementation includes the following features:

1. **Domain Driven Design** (DDD)
2. **Dependency Injection** (DI)
3. **Repository Pattern** (RP)
4. **Database Migration** (DM)
5. **API Documentation** (AD)

### Scenarios

- The system is a library with 3 roles: Director, Librarian, and Visitor.
- Everyone can find all the books in the library.
- registered visitors can borrow books, return books, and view their own borrowing records.
- librarians can update the information of books and have all the privileges of registered visitors.
- The director has full access to the librarians and can delete books. Deleting a book is a soft delete, so it does not affect the borrowing record.

## Packages

List third party packages with readme tables in this project

| Package                                       | Description                                                         |
| --------------------------------------------- | ------------------------------------------------------------------- |
| [gorm](https://gorm.io/)                      | The fantastic ORM library for Golang, aims to be developer friendly |
| [gorm/gen](https://gorm.io/gen/index.html)    | GORM GEN is a type-safe gorm code generator.                        |
| [gin](https://gin-gonic.com/)                 | Gin is a HTTP web framework written in Go (Golang)                  |
| [go-redis](https://github.com/redis/go-redis) | Redis client for golang                                             |
| [air](https://github.com/cosmtrek/air)        | Live reload for Go apps written in Go                               |

## Development steps

We suggest you to follow the steps below to develop with `Unix-like OS`, the whole project is developed and tested on `Ubuntu 22.04`.

### VSCode Plugins (Super Recommended)

`Visual Studio Code` is recommended as the IDE for this project, you can install the following extensions to make your development easier:

1. [Go](https://marketplace.visualstudio.com/items?itemName=golang.Go)
2. [Docker](https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-docker)
3. [OpenAPI (Swagger) Editor](https://marketplace.visualstudio.com/items?itemName=42Crunch.vscode-openapi)
4. [Caddyfile Support](https://marketplace.visualstudio.com/items?itemName=matthewpi.caddyfile-support)

### Third party tools (Optional)

The following tools are started by `docker` in the development environment, but you can install them on your machine if you want to:

- [DBeaver](https://dbeaver.io/download/) is recommended as the database management tool.
- [RedisInsight](https://redis.com/redis-enterprise/redis-insight/#insight-form) is recommended as the redis management tool.

### Setup

Here is the steps to start the development:

1. Clone the repo

   ```bash
   git clone https://github.com/alan890104/go-clean-architecture.git
   ```

2. Make sure you have installed `make`, `go` and `docker-compose / docker compose` on your machine
3. Copy `.env.example` to `.env.dev` and modify the environment variables if needed

   ```bash
   cp .env.example .env.dev
   ```

4. Run `make install` to install the tools
   > Run only once when you start the development or the tools required to be updated
5. Run `make dev-up` to start the development database and redis
6. Run `make dev-migrate` to migrate the database
7. Run `make generate` to generate the gorm/gen queries
8. Run `make serve` to start the development server with live reload
   1. Webserver will be listening on [localhost:8000](http://localhost:8000), you may change the port in `.env.dev`
   2. Swagger UI will be listening on [localhost:8080](http://localhost:8080)
   3. RedisInsight will be listening on [localhost:8001](http://localhost:8001)
   4. DBeaver will be listening on [localhost:8978](http://localhost:8978)
9. Run `make dev-down` to stop the development database and redis

## Iteration

We suggest the iteration steps for the development of this project as follows:

1. Have a API spec in openapi/grpc format (requires a lot of discussion)
2. Define the domain model
3. Define the repository interface
4. Define the usecase interface
5. Implement the controller
6. Implement the routes
7. Implement `cmd/app/app.go`
8. Implement the usecase
9. Implement the repository
