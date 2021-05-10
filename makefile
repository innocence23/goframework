.PHONY: mysql migrate-create createdb dropdb migrateup migratedown test server

USER = root
PWD = secret
PORT = 3307
DB = project

mysql57:
	docker run --name mysql8 -p 3306:3306  -e MYSQL_ROOT_PASSWORD=secret -d mysql:5.7

migrate-create:
	@echo "---Creating migration files---"
	migrate create -ext sql -dir ./migration -seq init_table

createdb:
	docker exec -it mysql57 mysql -u$(USER) -p$(PWD) -e 'create database $(DB)'

dropdb:
	docker exec -it mysql57 mysql -u$(USER) -p$(PWD) -e 'drop database $(DB)'

migrateup:
	migrate -path migration -database "mysql://$(USER):$(PWD)@tcp(localhost:$(PORT))/$(DB)" -verbose up

migratedown:
	migrate -path migration -database "mysql://$(USER):$(PWD)@tcp(localhost:$(PORT))/$(DB)" -verbose down

test:
	go test -v -cover ./...

server:
	go run main.go

