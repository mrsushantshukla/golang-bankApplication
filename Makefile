postgres:
	docker run --name=postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:16-alpine

deletepg:
	docker stop postgres16
	docker rm postgres16
createdb:
	docker exec -it postgres16 createdb --username=root --owner=root venus_bank
dropdb:
	docker exec -it postgres16 dropdb venus_bank
migrateup:
	migrate --path db/migrate -database "postgresql://root:root@localhost:5432/venus_bank?sslmode=disable" -verbose up 
migratedown:
	echo "Y" | migrate --path db/migrate -database "postgresql://root:root@localhost:5432/venus_bank?sslmode=disable" -verbose down 
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
.PHONY: postgres dropdb migrateup migratedown deletepg createdb sqlc test
