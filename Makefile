postgres:
	docker run --name ${DB_CONTAINER_NAME} -p 5432:5432 -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_OWNER) -d postgres:14-alpine

create_db:
	docker exec -it ${DB_CONTAINER_NAME} createdb --username=$(DB_USER) --owner=$(DB_OWNER) $(DB_NAME)

drop_db:
	docker exec -it ${DB_CONTAINER_NAME} dropdb $(DB_NAME)

migrate_up:
	migrate -path db/migration/ -database $(DB_SOURCE) -verbose up

migrate_down:
	migrate -path db/migration/ -database $(DB_SOURCE) -verbose down

sqlc:
	sqlc generate

check_install:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

generate_server:
	swagger generate model -f ./swagger.yaml -t ./api -A  stocksTracker

test:
	go test -v -race -short -cover ./api/...

mockgen:
	mockgen -package mockdb -destination db/mock/store.go github.com/TamerB/products-import-service/db/sqlc Store

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o "bin/products-import-service" "github.com/TamerB/products-import-service/api/cmd/products-import-service-server"

run:
	./bin/products-import-service

.PHONY: postgres createdb dropdb migrateup migratedown sqlc check_install generate_server test mockgen build run