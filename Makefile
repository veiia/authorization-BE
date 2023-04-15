.SILENT:

swagger:
	swag init -g cmd/auth/main.go

build: swagger
	docker-compose build auth-app # TODO: fixme

run: swagger
	docker-compose up -d auth-app # TODO: fixme

complex-run: swagger
	docker-compose up -d --build

test:

test-coverage:

create-migration:
	migrate create -ext sql -dir migrations -seq migration_name
migrate:
	docker run -v ./migrations:/migrations --network host migrate/migrate -path=/migrations -database "postgres://postgres:postgres@localhost:5436/postgres?sslmode=disable" up 1
migrate-down:
	docker run -v ./migrations:/migrations --network host migrate/migrate -path=/migrations -database "postgres://postgres:postgres@localhost:5436/postgres?sslmode=disable" down 1
migrate-drop:

psql:
	psql "sslmode=disable host=localhost dbname=postgres port=5436 user=postgres"
