# K8sCommerce - Local DB & Migrations

## Local Development

For local development you need a running database. To get a local database started quickly run:

```sh
docker-compose up
```

Docker will store the local Postgres volume here: 
```sh
db/.data
``` 

The `docker-compose up` creates a postgres database, applies the migrations and starts a local pgAdmin instance.


Check the `docker-compose` file for additional variables and paths.

## PGAdmin Access

Default username/password:

```sh
username: k8scommerce@example.com
password: postgres
```

## Goose Migrations

We use the [pressly/goose](https://github.com/pressly/goose) fork for db migrations. 
Set the env var `GOOSE_DBSTRING`
```sh
postgres://postgres:postgres@localhost:5432/k8scommerce?connect_timeout=180&sslmode=disable
```