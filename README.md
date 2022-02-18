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
First time? (*James Franco smile*)
```
go mod init backend
go mod tidy -compat=1.17
go run . -dummy -u 10
### number after -u is number of dummy users created
```
After that
```
go run . 
```