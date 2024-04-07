FROM golang:1.22.1-alpine3.19 as builder

WORKDIR /app

RUN apk add --no-cache make

COPY go.mod main.go ./
COPY . .
RUN go get

# --

FROM builder as build

WORKDIR /app

COPY --from=builder . ./

RUN make build

# --

FROM alpine:3.19

ENV ENV="prod"
EXPOSE 2000

WORKDIR /app

COPY --from=build ./app/bin/server ./

CMD ["./server"]
