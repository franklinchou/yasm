# yasm
yet another session management service

## Running locally

1. Starting redis

* `docker run --name redis-test -p 6379:6379 -d redis:latest`
* `docker start redis-test`

2. Running commands on the redis container

* `docker exec -it redis-test redis-cli`
