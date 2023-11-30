# nixpigweb

## Work In Progress

This is my first real project in Go. Maybe read about it [here](https://nixpig.dev)?

> No doubt there will be a bunch of areas for improvement. Feel free to let me know in the [issues](https://github.com/nixpig/nixpigweb/issues/new).

## Build

```shell
make tidy audit test build
```

## Deploy

### Provide `.env` file

```shell
SSL_CERT=
SSL_CERT_KEY=
SSL_CERT_DIRECTORY=

DATABASE_HOST=
DATABASE_PORT=
POSTGRES_USER=
POSTGRES_PASSWORD=
POSTGRES_DB=
DATABASE_MAX_CONNECTIONS=100
DATABASE_MAX_IDLE_CONNECTIONS=10
DATABASE_MAX_LIFETIME_CONNECTIONS=2

WEB_CONTEXT=
WEB_PORT=

API_CONTEXT=
API_PORT=

SECRET=

WAIT=5000000000

```

### Start containers

```shell
docker compose -p nixpig -f deploy/docker-compose.yml up -d

```

### Apply database migrations

```shell
migrate -path db/migrations  -database postgres://postgres:example_p4ssW0rd@localhost:5432/nixpigweb_?sslmode=disable up

```
