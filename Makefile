all: sqlc-gen style-build run

run:
	go run cmd/app/main.go

sqlc-gen:
	sqlc generate

up: compose-up migrate-up

compose-up:
	docker-compose up -d

compose-down:
	docker-compose down

migrate-up:
	go run cmd/migrate/main.go -run=up

migrate-down:
	go run cmd/migrate/main.go -run=down

style-build:
	npx tailwindcss -i ./assets/css/src.css -o ./assets/css/dist.css
