include .env
export

DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

.PHONY: migrate-create migrate-up migrate-down

migrate-create:
	migrate create -ext sql -dir deployments/migrations -seq $(name)

migrate-up:
	migrate -path deployments/migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path deployments/migrations -database "$(DB_URL)" down 1

migrate-force:
	migrate -path deployments/migrations -database "$(DB_URL)" force $(version)