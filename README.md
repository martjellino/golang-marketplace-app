# golang-marketplace-app

## Prerequisites
1. Installed ```golang-migrate-cli```, here is the how to https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

## How to run
1. Run ```go mod tidy``` to install & remote unecessary dependencies
2. ...

## Migration guide
1. To execute migration
```
migrate -database "postgres://{username}:{password}@{host}:{port}/{dbname}?sslmode=disable" -path migrations up
```

2. To rollback migration
```
migrate -database "postgres://{username}:{password}@{host}:{port}/{dbname}?sslmode=disable" -path migrations down
```

3. Example
```
migrate -database "postgres://postgres:P4ssW0rd@localhost:5434/marketplace_db?sslmode=disable" -path migrations up
```

4. If you want to add postgres docker for local development
```
docker pull postgres
docker run --name marketplace-app -e POSTGRES_PASSWORD="P4ssW0rd" -e POSTGRES_DB="marketplace_db" -d -p 5434:5432 postgres
```
default username is 'postgres'
