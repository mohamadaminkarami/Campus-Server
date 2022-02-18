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
how to run
```
go mod init backend
go mod tidy -compat=1.17
go run . 
### number after -u is number of dummy users created
```