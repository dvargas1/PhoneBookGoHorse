FROM golang:1.22.4-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

WORKDIR /app/cmd/phonebook

RUN go build -o /app/cmd/phonebook .

EXPOSE 7331
ENTRYPOINT [ "/app/cmd/phonebook/phonebook" ]

