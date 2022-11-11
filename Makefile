.SILENT:

build:
	docker-compose build auth-app

run:
	docker-compose up -d auth-app

complex-run:
	docker-compose up -d --build

test:

test-coverage:

create-migration:

migrate:

migrate-down:

migrate-drop:

swagger:
	swag init -g cmd/auth/main.go

