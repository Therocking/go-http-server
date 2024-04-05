
build:
	go build -o bin/server .

dev:
	export ENV=dev
	go run .

run:
	bin/server
