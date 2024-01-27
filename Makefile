.PHONY: clean critic security lint test build run

APP_NAME = shipment
BUILD_DIR = $(PWD)/build
MIGRATIONS_FOLDER = $(PWD)/infrastructure/pkg/sql
DATABASE_URL=postgresql://myuser:password@localhost:5432/postgres?sslmode=disable

clean:
	rm -rf ./build

build:
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" up

migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down

migrate.force:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" force $(version)

docker.run: docker.postgres docker.redis docker.kafka migrate.up

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
