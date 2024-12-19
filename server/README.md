## Feature Flags

Server side application of feature flags

## Installation

```bash
$ go mod download
```

## Running with Docker

```bash
# build & run
$ docker-compose up --build
```

```bash
# connect to postgres
$ psql -h localhost -U postgres -d feature_flags
```

## Run migrations manually

```bash
# up
$ goose -dir [migrations-directory] [driver] ["db-url"] up
```

```bash
# down
$ goose -dir [migrations-directory] [driver] ["db-url"] down
```
