# HTTP go server
Made using `fiber` frameword and `grom` ORM.

## Run with docker-compose
The docker-compose file create the image of the go app and the postgres db.
```bash
docker compose up
```

## Run app (local)
1. Clone `.env.template` file to `.env` and config envs.
2. Config local env `ENV` and assing the valid `dev`. ex:
- linux
```bash
export ENV=dev
```
3. Run postgres db:
```bash
docker run \
--name db \
-dp 5432:5432 \
-e POSTGRES_PASSWORD=password \
postgres:16.2-alpine3.19
```
4. Run `make dev`, to run the app. If you dont have `make` instaled, run this instaed: `go run .`.

## Run image
```bash
docker run \
--name go_server \
-dp 2000:2000 \
-e DB_USER=postgres \
-e DB_PASS=password \
-e DB_HOST=db \
-e DB_PORT=5432 \
-e DB_NAME=users_products \
go_server:latest
```
