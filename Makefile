.SILENT:

swagger:
	swag init -g cmd/auth/main.go

build: swagger
	docker-compose build auth-app

run: swagger
	docker-compose up -d auth-app

complex-run: swagger
	docker-compose up -d --build

test:

test-coverage:

create-migration:

migrate:
	docker run -v /home/syth0le/coding/diplom/authorization-BE/migrations:/migrations --network host migrate/migrate -path=/migrations -database "postgres://postgres:postgres@localhost:5436/postgres?sslmode=disable" up 1
migrate-down:
	docker run -v /home/syth0le/coding/diplom/authorization-BE/migrations:/migrations --network host migrate/migrate -path=/migrations -database "postgres://postgres:postgres@localhost:5436/postgres?sslmode=disable" down 1
migrate-drop:


