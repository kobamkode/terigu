air: style-build build-air

run:
	go run cmd/app/main.go

build-air:
	go build -o tmp/main cmd/app/main.go

sqlc-gen:
	sqlc generate

compose-up:
	docker-compose up -d

compose-down:
	docker-compose down

migrate-up:
	go run cmd/migrate/main.go -run=up

migrate-down:
	go run cmd/migrate/main.go -run=down

style-build:
	npx tailwindcss -i ./views/assets/css/src.css -o ./views/assets/css/dist.css
