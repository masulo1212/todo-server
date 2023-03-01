postgres:
	docker run --name DB -p 5432:5432  -e POSTGRES_USER=root -e POSTGRES_PASSWORD=0000 -e GIN_MODE=release -d postgres

createdb:
	docker exec -it DB createdb --username=root --owner=root todoDB
dropdb:
	docker exec -it DB dropdb todoDB
migrateup:
	migrate -path broker/db/migration -database "postgresql://root:0000@localhost:5432/todoDB?sslmode=disable" -verbose up
migrateup1:
	migrate -path broker/db/migration -database "postgresql://root:0000@localhost:5432/todoDB?sslmode=disable" -verbose up 1
migratedown:
	migrate -path broker/db/migration -database "postgresql://root:0000@localhost:5432/todoDB?sslmode=disable" -verbose down
migratedown1:
	migrate -path broker/db/migration -database "postgresql://root:0000@localhost:5432/todoDB?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate


BROKER_BINARY=brokerApp
BROKER_BINARY_amd64=brokerApp_amd64
DISPATCH_BINARY=dispatchApp


up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

down:
	@echo "Stopping docker images"
	docker-compose down
	@echo "Docker images stop!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_broker
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## build_broker: builds the broker binary as a linux executable
build_broker:
	@echo "Building broker binary..."
	cd ./broker && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} .
	@echo "Done!"

build_broker_amd64:
	@echo "Building broker binary..."
	cd ./broker && env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o ${BROKER_BINARY_amd64} .
	@echo "Done!"

## build_broker: builds the broker binary as a linux executable
build_dispatch:
	@echo "Building broker binary..."
	cd ./broker && env GOOS=linux CGO_ENABLED=0 go build -o ${DISPATCH_BINARY} .
	@echo "Done!"