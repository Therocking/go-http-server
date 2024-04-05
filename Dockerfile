FROM golang:1.21.9-alpine3.19

ENV ENV="prod"
ENV PORT=2000
EXPOSE 2000

WORKDIR /app

COPY go.mod main.go go.sum ./
COPY . .

RUN go build -o bin .

CMD ["./bin"]
