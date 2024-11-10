# Simple API 

- Go ( gin )
- MySQL
- sqlc [document](https://docs.sqlc.dev/en/stable/tutorials/getting-started-mysql.html)
- Docker
- database migration [goose](https://github.com/pressly/goose)

## Env Example

change config env file  ->  main.go , Makefile

```bash
APP_PORT=

CONTAINER_NAME=

MYSQL_ROOT_PASSWORD=
MYSQL_DBNAME=
MYSQL_USER=
MYSQL_PASSWORD=
MYSQL_DSN=
```

## Usage

```bash
# docker compose up
make docker-up

# docker compose down
make docker-down

# goose database migration up
make migrate-up

# goose database migration down
make migrate-down

# sqlc generate
make sqlc
```
