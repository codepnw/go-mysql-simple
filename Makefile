include dev.env

docker-up:
	docker compose --env-file dev.env up --build -d

docker-down:
	docker compose down

migrate-up: 
	cd internal/database/schema && goose mysql ${MYSQL_DSN} up

migrate-down:
	cd internal/database/schema && goose mysql ${MYSQL_DSN} down

sqlc:
	sqlc generate