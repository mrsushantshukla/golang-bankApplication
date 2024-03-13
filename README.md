# golang-bankApplication
A bank application made in Go

Steps followed:
1. Install Postgres Image (postgres:16-alpine)
2. docker run --name=postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:16-alpine
3. brew install golang-migrate
4. Create a new folder. Inside it create db/migrate folder.
5. migrate create -ext sql -dir db/migrate -seq init_schema
This will generate an empty migration file which we can use to either start a new connection (up file) or in case of downgrading (down file)
6. Copy the script generated from dbdiagram.io to the up file.
7. In down file, write sql code to drop all tables.
8. Inside docker (postgres) -> createdb --username=root --owner=root venus_bank
9. Make Makefile with setup cmds
10. migrate --path db/migrate -database "postgresql://root:root@localhost:5432/venus_bank?sslmode=disable" -verbose up 
11. Install sqlc library which will be used to write sql queries -> brew install sqlc
12. Initialize the repo with the command -> sqlc init
13. Postgres driver -  go get github.com/lib/pq
14. Run tests in db/sqlc/ - go test
15. go get github.com/stretchr/testify    - to test
