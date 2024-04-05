FROM --platform=$BUILDPLATFORM golang:1.21.9-alpine3.19

ENV ENV="prod"
EXPOSE 2000

WORKDIR /app

RUN apk add make

COPY go.mod main.go ./
COPY . .
RUN go get

RUN make build

CMD ["make", "run"]
