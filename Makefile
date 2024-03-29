.PHONY: clean critic security lint test build run

APP_NAME = commerce-shipment
BUILD_DIR = $(PWD)/build
MIGRATIONS_FOLDER = $(PWD)/internal/adapter/storage/postgres/migrations
DATABASE_URL=postgres://user:password@localhost:5432/db?sslmode=disable
VERSION=1

clean:
	rm -rf ./build

build:
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" -verbose up

migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" -verbose down

migrate.force:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" force $(VERSION)

docker.build:
	docker build -t $(APP_NAME) .

docker.run: docker.postgres docker.redis

docker.postgres:
	docker run --rm -d \
		--name my_postgres_container\
		-e POSTGRES_USER=user\
		-e POSTGRES_PASSWORD=password\
		-e POSTGRES_DB=db\
		-p 5432:5432\
		postgres

docker.redis:
	docker run --rm -d \
		--name my_redis_container \
          -e REDIS_PASSWORD=password \
          -e REDIS_DB=0 \
          -p 6379:6379 \
          redis

docker.kafka:
	docker run --rm -d \
		--name cgapp-kafka \
		-p 9092:9092 \
		-e KAFKA_ZOOKEEPER_CONNECT=localhost:2181 \
		-e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092 \
		-e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 \
		bitnami/kafka:latest

docker.stop: docker.stop.postgres docker.stop.redis docker.stop.kafka migrate.down

docker.stop.fiber:
	docker stop cgapp-fiber

docker.stop.postgres:
	docker stop cgapp-postgres

docker.stop.redis:
	docker stop cgapp-redis

docker.stop.kafka:
	docker stop cgapp-kafka
