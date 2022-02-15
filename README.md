# Campus-Server
postgres connection .env
```
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgrespass
POSTGRES_DB_NAME=postgres
POSTGRES_TIMEZONE=Asia/Tehran
POSTGRES_SSL_MODE=disable
JWT_SECRET=VERYSECRET
```

```
go mod init backend
go mod tidy
go run .
```