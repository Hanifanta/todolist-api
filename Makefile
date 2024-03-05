createdb:
	sudo docker exec -it mysql-docker mysql -u root -proot -e "CREATE DATABASE IF NOT EXISTS todolist"

dropdb:
	sudo docker exec -it mysql-docker mysql -u root -proot -e "DROP DATABASE IF EXISTS todolist"

migratecreate:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrateup:
	migrate -path db/migrations -database "mysql://root:root@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local" -verbose up

migratedown:
	migrate -path db/migrations -database "mysql://root:root@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local" -verbose down 1

migratefix:
	migrate -path db/migrations -database "mysql://root:root@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local" -verbose force $(version)

test:
	go test -v -cover ./...

PHONY: createdb dropdb migrate_create migrate_up migrate_down test migrate_fix