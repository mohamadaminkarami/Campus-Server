FROM golang:alpine

WORKDIR /go/src/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /server

RUN chmod +x wait-for-it.sh

EXPOSE 8080

CMD [ "/server" ]