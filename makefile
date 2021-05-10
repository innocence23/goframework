.PHONY: mysql migrate-create createdb dropdb migrateup migratedown genemodel test server

USER = root
PWD = root
PORT = 3306
DB = project

mysql57:
	docker run --name mysql8 -p 3306:3306  -e MYSQL_ROOT_PASSWORD=$(PWD) -d mysql:5.7

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

genemodel:
	model-generator -u=$(USER) -p=$(PWD) -d=$(DB)  -dir=./model

test:
	go test -v -cover ./...

server:
	go run main.go

