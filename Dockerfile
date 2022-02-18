FROM golang:alpine

WORKDIR /go/src/app

RUN apk --update add postgresql-client

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /server

RUN chmod +x wait-for-postgres.sh

EXPOSE 8080

CMD [ "/server" ]