postgres_init:
    docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:16-alpine

postgres:
	docker exec -it postgres15 psql

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root db

dropdb:
	docker exec -it postgres15 dropdb db

.PHONY: postgres_init postgres createdb dropdb